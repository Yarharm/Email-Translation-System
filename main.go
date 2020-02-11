package main

import (
	"fmt"
	"log"
)

func main() {
	client := authenticateClient()

	unreadMessages := getUnreadMessages(client)

	for _, msg := range unreadMessages {
		responseTexts, err := translateMessages(msg.snippet, TranslationCode)
		if err != nil {
			log.Fatalf("Translation failed: %v", err)
			return
		}
		for _, translatedText := range responseTexts {
			fmt.Println(translatedText)
		}
	}
	return
}
