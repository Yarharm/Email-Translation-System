package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// TranslationResponse struct
type translationResponse struct {
	Code int      `json:"code"`
	Lang string   `json:"lang"`
	Text []string `json:"text"`
}

// TranslateToken struct
type translateToken struct {
	Token string `json:"token"`
}

func translateMessages(snippet string, translationCode string) ([]string, error) {
	key, err := readTranslationToken()
	if err != nil {
		log.Fatalf("Unable to read translation token: %v", err)
		return nil, err
	}

	responseTexts, err := sendOverTheWire(snippet, key, translationCode)
	if err != nil {
		log.Fatalf("Failure to retrieve translated text: %v", err)
		return nil, err
	}
	return responseTexts, nil
}

func sendOverTheWire(textValue string, keyValue string, langValue string) ([]string, error) {
	resp, err := http.PostForm(YandexTranslateDomain,
		url.Values{"key": {keyValue}, "text": {textValue}, "lang": {langValue}})
	if err != nil {
		log.Fatalf("Unable to access Translate API: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	payload := translationResponse{}
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Fatalf("Unable to parse translation response: %v", err)
		return nil, err
	}
	return payload.Text, nil
}

func readTranslationToken() (string, error) {
	tokFile := TranslationTokenFile
	var payload translateToken

	tokenJSON, err := os.Open(tokFile)
	if err != nil {
		log.Fatalf("Unable to find translation token file: %v", err)
		return "", nil
	}

	tok, _ := ioutil.ReadAll(tokenJSON)
	json.Unmarshal(tok, &payload)

	return payload.Token, nil
}
