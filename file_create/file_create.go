package file_create

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func JSON_parse(name string, jb []byte) error {
	absPath, err := filepath.Abs("./text_analysis/json_files")
	if err != nil {
		return errors.New(fmt.Sprintf("помилка отримання абсолютного шляху до файлу: %v", err))
	}
	f, err := os.OpenFile(filepath.Join(absPath, fmt.Sprintf("%v.json", name)), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return errors.New(fmt.Sprintf("помилка опрацювання json файла: %v", err))
	}
	defer f.Close()
	fmt.Println(string(jb))
	_, err = f.Write(jb)
	if err != nil {
		return errors.New(fmt.Sprintf("Помилка вписання в файл json: %v", err))
	}

	// Move file pointer to the beginning before reading
	if _, err = f.Seek(0, 0); err != nil {
		return errors.New(fmt.Sprintf("Помилка переміщення покажчика файлу: %v", err))
	}

	readBytes, err := io.ReadAll(f)
	if err != nil {
		return errors.New(fmt.Sprintf("Помилка прочитання файла: %v", err))
	}
	fmt.Println(string(readBytes))
	return nil
}
