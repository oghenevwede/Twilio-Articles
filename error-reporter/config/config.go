package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	TeamEmails   map[string]string `json:"team_emails"`
	DefaultEmail string            `json:"default_email"`
}

// LoadConfig loads the configuration from the config.json file
func LoadConfig() (Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
