package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) {
	c.CurrentUserName = username

	filePath, err := getFilePath()
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		log.Fatal(err)
	}

}
func Read() *Config {
	filePath, err := getFilePath()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("could not read %s file: %v\n", configFileName, err)
	}

	cfg := Config{}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal %s config: %v\n", configFileName, err)
	}

	return &cfg
}

func getFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user's home directory: %v", err)
	}

	filePath := filepath.Join(homeDir, configFileName)

	return filePath, nil
}
