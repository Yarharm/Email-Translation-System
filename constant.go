package main

const (

	// AuthenticationTokenFile path to authentication token
	AuthenticationTokenFile string = "token.json"

	// Credentials path
	Credentials string = "credentials.json"

	// TranslationTokenFile path to Yandex translate token
	TranslationTokenFile string = "translation_token.json"

	// UnreadMessagesQuery to fetch unread messages from the inbox
	UnreadMessagesQuery string = "in:inbox is:unread category:primary"

	// TargetUserID for gmail account
	TargetUserID string = "me"

	// YandexTranslateDomain translation API domain
	YandexTranslateDomain string = "https://translate.yandex.net/api/v1.5/tr.json/translate"

	// TranslationCode specifies translation language
	TranslationCode string = "en-ru"
)
