package main

import (
	"encoding/json"
	"os"
)

// Configuration contains values from conf.json
type Configuration struct {
	Lyft struct {
		ClientID     string `json:"ClientId"`
		ClientSecret string `json:"ClientSecret"`
	} `json:"Lyft"`
	MapQuest struct {
		ConsumerKey string `json:"ConsumerKey"`
	}
	Uber struct {
		ServerToken string `json:"ServerToken"`
	} `json:"Uber"`
}

// LoadConfiguration loads config data from conf.json
func LoadConfiguration() Configuration {
	var configuration Configuration
	configFile, err := os.Open("conf.json")
	if err != nil {
		panic(err.Error())
	}

	parser := json.NewDecoder(configFile)
	parser.Decode(&configuration)

	configFile.Close()

	return configuration
}
