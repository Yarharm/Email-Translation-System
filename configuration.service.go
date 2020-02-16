package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type configService struct {
	TranslationRecipient string `json:"translationRecipient"`
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
