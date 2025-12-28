package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to fetch user's home directory: %v\n", err)
	}

	filePath := filepath.Join(homeDir, ".gatorconfig.json")
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("could not read gatorconfig.json file: %v\n", err)
	}

	cfg := Config{}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal gatorconfig.json data: %v\n", err)
	}

	return &cfg
}
