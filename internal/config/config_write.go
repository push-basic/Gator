package config

import (
	"encoding/json"
	"fmt"
	"os"
)

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
