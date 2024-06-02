package text_analysis

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func JsonPrepoc(filename string) error {
	root, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return err
	}

	jsonFilePath := filepath.Join(root, "text_analysis", "json_files", fmt.Sprintf("%s.json", filename))
	f, err := os.Open(jsonFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	jsonFile, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return err
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(jsonFile, &obj); err != nil {
		fmt.Printf("Error unmarshalling JSON file: %v\n", err)
		return err
	}
	var texts []string
	if messages, ok := obj["messages"].([]interface{}); ok {
		for _, message := range messages {
			if msgMap, ok := message.(map[string]interface{}); ok {
				if text, ok := msgMap["text"].(string); ok {
					texts = append(texts, text)
				}
			}
		}
	}
	combinedText := strings.Join(texts, "\n")

	normalizedText := normalizeWhitespace(combinedText)

	txtFilePath := filepath.Join(root, "text_analysis", "txt_files", fmt.Sprintf("%s.txt", filename))
	if err := os.WriteFile(txtFilePath, []byte(strings.ToLower(normalizedText)), 0644); err != nil {
		fmt.Printf("Error writing text file: %v\n", err)
		return err
	}

	fmt.Println("Text processing completed successfully.")
	return nil
}

// normalizeWhitespace replaces non-word characters and extra spaces with a single space
func normalizeWhitespace(input string) string {
	reNonWord := regexp.MustCompile(`\W+`)
	reExtraSpaces := regexp.MustCompile(`\s+`)

	input = reNonWord.ReplaceAllString(input, " ")
	return reExtraSpaces.ReplaceAllString(input, " ")
}

//
//// processText processes the text field from the JSON message
//func processText(text interface{}) string {
//	texts := ""
//	switch t := text.(type) {
//	case string:
//		texts += "\n" + t
//	case []interface{}:
//		for _, item := range t {
//			switch i := item.(type) {
//			case string:
//				texts += "\n" + i
//			case map[string]interface{}:
//				if subText, ok := i["text"].(string); ok {
//					texts += "\n" + subText
//				}
//			}
//		}
//	}
//	return texts
//}
//
//// normalizeWhitespace replaces non-word characters and extra spaces with a single space
//func normalizeWhitespace(input string) string {
//	reNonWord := regexp.MustCompile(`\W+`)
//	reExtraSpaces := regexp.MustCompile(`\s+`)
//
//	input = reNonWord.ReplaceAllString(input, " ")
//	return reExtraSpaces.ReplaceAllString(input, " ")
//}
