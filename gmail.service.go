package main

import (
	"encoding/base64"
	"log"

	gmail "google.golang.org/api/gmail/v1"
)

type message struct {
	gmailID string
	snippet string
}

func sendTranslatedText(client *gmail.Service, message string, config *configService) {
	msg := buildResponseMessage(message, config)
	_, err := client.Users.Messages.Send(TargetUserID, &msg).Do()
	if err != nil {
		log.Fatalf("Unable to send translated text %v", err)
	}
}

func buildResponseMessage(body string, config *configService) gmail.Message {
	var message gmail.Message
	messageStr := []byte("From: 'me'\r\n" +
		"To: " + config.TranslationRecipient + "\r\n" +
		"Subject: " + config.Subject + " \r\n" +
		"\r\n" + body)
	message.Raw = base64.RawURLEncoding.EncodeToString((messageStr))
	return message
}

func getUnreadMessages(client *gmail.Service) []message {
	return listMessages(client, TargetUserID, UnreadMessagesQuery)
}

func listMessages(service *gmail.Service, userID string, query string) []message {
	messages := []message{}
	pageToken := ""
	request := service.Users.Messages.List(userID).Q(query)
	for {
		if pageToken != "" {
			request.PageToken(pageToken)
		}

		req, err := request.Do()
		if err != nil {
			log.Fatalf("Unable to retrieve messages: %v", err)
		}

		log.Printf("Processing %v messages...\n", len(req.Messages))
		for _, m := range req.Messages {
			msg, err := service.Users.Messages.Get(userID, m.Id).Do()
			if err != nil {
				log.Fatalf("Unable to retrieve message %v: %v", m.Id, err)
			}
			messages = append(messages, message{
				gmailID: msg.Id,
				snippet: msg.Snippet,
			})
		}

		if req.NextPageToken == "" {
			break
		}
		pageToken = req.NextPageToken
	}

	return messages
}
