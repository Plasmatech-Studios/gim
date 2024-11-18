package config

import (
	"encoding/json"
	"fmt"
	"gim/internal/utils"
	"os"
	"path/filepath"
)

type Config struct {
    Aliases map[string]string `json:"aliases"`
    Using   string            `json:"using"` // Tracks the current alias
}


// LoadConfig loads the configuration file from disk.
func LoadConfig() (*Config, error) {
	configPath := GetConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createDefaultConfig(configPath); err != nil {
			return nil, fmt.Errorf("could not create default config: %w", err)
		}
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not open config: %w", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("could not decode config: %w", err)
	}
	return &config, nil
}

// SaveConfig saves the configuration to disk.
func SaveConfig(cfg *Config) error {
	configPath := GetConfigPath()
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("could not save config: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cfg)
}

// GetConfigPath returns the path to the configuration file.
func GetConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Could not determine home directory")
	}
	return filepath.Join(home, ".gim", "config.json")
}

// createDefaultConfig initializes a default configuration.
func createDefaultConfig(configPath string) error {
	sshDir := utils.GetSSHDir()
	files, err := os.ReadDir(sshDir)
	if err != nil {
		return fmt.Errorf("could not read .ssh directory: %w", err)
	}

	aliases := make(map[string]string)
	for _, file := range files {
		if utils.IsPrivateKey(file.Name()) {
			alias := file.Name()
			aliases[alias] = filepath.Join(sshDir, alias)
		}
	}

	if len(aliases) == 0 {
		aliases["example"] = filepath.Join(sshDir, "id_example")
	}

	defaultConfig := Config{Aliases: aliases}

	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("could not create config directory: %w", err)
	}

	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("could not create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(defaultConfig)
}
