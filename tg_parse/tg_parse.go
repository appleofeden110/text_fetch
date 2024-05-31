package tg_parse

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	teleg "github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"io"
	"log"
	"os"
	"strings"
)

var (
	phone         string
	chat_username string
	password      string
)

func TelegramParse(ctx context.Context, api_id int, api_hash string) error {
	var err error

	if err != nil {
		return err
	}
	api_hash = os.Getenv("API_APP_HASH")
	client := teleg.NewClient(api_id, api_hash, teleg.Options{})
	if err := client.Run(ctx, func(ctx context.Context) error {
		defer func() {
			if _, err := client.API().AuthLogOut(ctx); err != nil {
				log.Printf("Failed to log out: %v", err)
			} else {
				fmt.Println("Logged out successfully")
			}
		}()
		//checks if the password needed. DO NOT CHANGE errCodeVAR NAME
		_, errCode := auth.CodeOnly(phone, auth.CodeAuthenticatorFunc(codeAsk)).Password(ctx)
		tokMap, err := createAuthToken(errCode)
		if err != nil {
			return errors.New(fmt.Sprintf("помилка з auth файлом: %v\n", err))
		}
		if tokMap != nil {
			phone = tokMap["phone"]
			password = tokMap["password"]
		}
		fmt.Println(phone)
		if errors.Is(errCode, auth.ErrPasswordNotProvided) {
			if err != nil {
				return err
			}
			err = auth.NewFlow(
				auth.Constant(phone, password, auth.CodeAuthenticatorFunc(codeAsk)),
				auth.SendCodeOptions{},
			).Run(ctx, client.Auth())
			if err != nil {
				return err
			}
		} else {
			err = auth.NewFlow(
				auth.CodeOnly(string(phone), auth.CodeAuthenticatorFunc(codeAsk)),
				auth.SendCodeOptions{},
			).Run(ctx, client.Auth())
			if err != nil {
				return err
			}
		}
		fmt.Print("Введіть Username чату з якого хочете взяти повідомлення:")
		_, err = fmt.Scanln(&chat_username)
		msgs, err := MessageFetch(ctx, client.API(), chat_username)
		if err != nil {
			return err
		}
		for _, msg := range msgs {
			fmt.Println(msg.Message)
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func codeAsk(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
	fmt.Print("Введіть код прийшовший на телеграм з вище наведеним номером телефону для авторизації:")
	code, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Problem with reading a string")
		return "", err
	}
	code = strings.ReplaceAll(code, "\n", "")
	return code, nil
}

func MessageFetch(ctx context.Context, client *tg.Client, username string) ([]*tg.Message, error) {
	// Search for a public chat by username
	chat, err := client.ContactsResolveUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	channel := chat.Chats[0].(*tg.Channel)

	var limit int

	fmt.Print("Ліміт повідомлень? (число або нажміть Enter для всіх):")
	_, err = fmt.Scanln(&limit)
	peer := &tg.InputPeerChannel{channel.ID, channel.AccessHash}
	dgs, err := client.MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{Peer: peer, Limit: limit})
	if err != nil {
		return nil, err
	}
	messageClass, ok := dgs.(*tg.MessagesChannelMessages)
	if !ok {
		return nil, errors.New(fmt.Sprintf("unexpected msg class %T", dgs))
	}

	messages := make([]*tg.Message, 0, len(messageClass.Messages))
	for _, msg := range messageClass.Messages {
		if message, ok := msg.(*tg.Message); ok {
			messages = append(messages, message)
		}
	}
	return messages, nil
}

func createAuthToken(e error) (map[string]string, error) {
	//checking if the file exists, if not, creating new file with data in it
	f, err := os.OpenFile("auth_token.json", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("помилка з відкриттям файлу для авторизація: %v\n", err))
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("помилка прочитання файла: %v\n", err))
	}
	if string(b) == "" {
		fmt.Print("Введіть свій номер телефону для входу в телеграм:")
		_, err = fmt.Scanln(&phone)
		if errors.Is(e, auth.ErrPasswordNotProvided) {
			fmt.Print("Введіть пароль (тільки якщо у вас включена 2-factor auth):")
			_, err = fmt.Scanln(&password)
		}
		n, err := f.Write([]byte(fmt.Sprintf(`{"phone":"%v","password":"%v"}`, phone, password)))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("помилка з написанням файлу для авторизації: %v\n", err))
		}
		fmt.Println(n)
	}
	authCred := make(map[string]string)
	err = json.Unmarshal(b, &authCred)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("помилка з парсингом json файлу: %v\n", err))
	}
	return authCred, nil
}
