package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type configService struct {
	TranslationLanguage  string `json:"translationLanguage"`
	TranslationRecipient string `json:"translationRecipient"`
	Subject              string `json:"subject"`
}

// Retrieves a configuration from a local file.
func spawnConfigService() (*configService, error) {
	configJSON, err := os.Open(ConfigurationFile)
	var configService configService
	if err != nil {
		return nil, err
	}
	defer configJSON.Close()
	config, _ := ioutil.ReadAll(configJSON)
	json.Unmarshal(config, &configService)
	return &configService, nil
}
