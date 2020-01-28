package main

import (
	"log"
	"net/http"

	gmail "google.golang.org/api/gmail/v1"
)

type message struct {
	gmailID string
	snippet string
}

func getUnreadMessages(client *http.Client) []message {
	service, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to create Gmail service: %v", err)
	}

	return listMessages(service, TargetUserID, UnreadMessagesQuery)
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
