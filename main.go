package main

import (
	"fmt"
	"log"
)

func main() {
	config, error := spawnConfigService()
	if error != nil {
		log.Fatalf("Could read configuration file: %v", error)
	}
	client := authenticateClient()

	unreadMessages := getUnreadMessages(client)

	for _, msg := range unreadMessages {
		responseTexts, err := translateMessages(msg.snippet, TranslationCode)
		if err != nil {
			log.Fatalf("Translation failed: %v", err)
			return
		}
		for _, translatedText := range responseTexts {
			fmt.Println(config)
			fmt.Println(translatedText)
		}
	}
	return
}
