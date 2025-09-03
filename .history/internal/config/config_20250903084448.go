package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// Config represents the structure of the JSON config file.
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

// Read loads the config from ~/.gatorconfig.json
func Read() (Config, error) {
	var cfg Config

	path, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

// SetUser updates the current_user_name and writes back to file
func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return write(*c)
}

// -------------------- Helpers --------------------

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find home dir: %w", err)
	}
	return filepath.Join(home, configFileName), nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
