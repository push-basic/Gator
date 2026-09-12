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
	Url  string `json:"url"`
	Name string `json:"name"`
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
