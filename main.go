package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"text_fetch/file_create"
	"text_fetch/text_analysis"
	"text_fetch/tg_parse"
	"text_fetch/yt_parse"
)

var (
	tg_api_id   int
	tg_api_hash string
	ctx         context.Context
)

func check(err error, msg ...string) {
	if err != nil {
		log.Fatalf("%v: %v", msg, err)
		return
	}
}

func main() {
	var err error
	ctx = context.Background()
	err = godotenv.Load()
	check(err, ".env")
	tg_api_hash = os.Getenv("API_APP_HASH")
	tg_api_id, err = strconv.Atoi(os.Getenv("API_APP_ID"))
	if err != nil {
		log.Fatalf("A problem with string convertation: %v\n", err)
		return
	}

	fmt.Print("Яка із платформ вас цікавить телеграм чи ютуб?(T/Y):")
	var choice string
	_, err = fmt.Scanln(&choice)
	check(err, "choice read error")
	switch choice {
	case "T":
		msgs, err := tg_parse.TelegramParse(ctx, tg_api_id, tg_api_hash)
		check(err, "tg_parse")
		jBytes, errJ := tg_parse.MarshalJSON(msgs)
		check(errJ, "MARSHAL JSON")
		err = file_create.JSON_parse("test_tg", jBytes)
		check(err, "json_parse tg")
		err = text_analysis.JsonPrepoc("test_tg")
		check(err, "json_preproc tg")
		err = text_analysis.TextAnalysis("test_tg")
		check(err, "text_analysis tg")
		break
	case "Y":
		jsonBytes, err := yt_parse.YoutubeParse(ctx)
		check(err, "youtube parse")
		err = file_create.JSON_parse("test_yt", jsonBytes)
		check(err, "json_parse")
		break
	default:
		check(errors.New("Неправильний вибір, напишіть T або Y для вибора між Телеграмом та Ютубом"), "mistype of type of parse")
	}
}
