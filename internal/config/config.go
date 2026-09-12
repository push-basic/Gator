package config

import (
	"encoding/json"
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

func Read() (Config, error) {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("getting config path: %w", err)
	}

	jsonData, err := os.ReadFile(cfgPath)
	if err != nil {
		return Config{}, fmt.Errorf("reading config file: %w", err)
	}

	var config Config

	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return Config{}, fmt.Errorf("unmarshalling json: %w", err)
	}

	return config, nil
}

func (cfg *Config) SetUser(param string) error {
	cfg.Name = param
	return writeConfig(*cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting home directory: %w", err)
	}
	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}

func writeConfig(cfg Config) error {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("getting config path: %w", err)
	}

	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("mashalling config: %w", err)
	}

	err = os.WriteFile(cfgPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("writing config file: %w", err)
	}

	return nil
}
