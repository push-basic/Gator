package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	URL  string `json:"db_url"`
	Name string `json:"current_user_name"`
}

func (c *Config) SetUser(param string) error {
	c.Name = param
	return writeConfig(*c)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting home directory: %w", err)
	}
	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}
