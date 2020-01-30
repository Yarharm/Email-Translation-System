package main

import (
	"fmt"
)

func main() {
	client := authenticateClient()

	unreadMessages := getUnreadMessages(client)

	for _, msg := range unreadMessages {
		fmt.Printf("%s --> %s\n", msg.gmailID, msg.snippet)
	}
}
