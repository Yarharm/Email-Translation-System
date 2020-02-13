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
		responseTexts, err := translateMessages(msg.snippet, config.TranslationLanguage)
		if err != nil {
			log.Fatalf("Translation failed: %v", err)
			return
		}
		for _, translatedText := range responseTexts {
			sendTranslatedText(client, translatedText, config)
			fmt.Println("Successfully sent translated email")
		}
	}
}
