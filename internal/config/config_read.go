package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadConfig() (Config, error) {
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
