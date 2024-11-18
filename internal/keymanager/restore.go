package keymanager

import (
	"fmt"
	"gim/internal/config"
	"gim/internal/utils"
	"path/filepath"
)

func RestoreKey(alias string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if _, exists := cfg.Aliases[alias]; exists {
		return fmt.Errorf("alias '%s' already exists in config", alias)
	}

	// Automatically find the orphaned key path
	keyPath := filepath.Join(utils.GetSSHDir(), alias)
	if !utils.IsFileExists(keyPath) {
		return fmt.Errorf("no orphaned key found for alias '%s'", alias)
	}

	// Add alias to config
	cfg.Aliases[alias] = keyPath
	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to save updated config: %w", err)
	}

	fmt.Printf("Alias '%s' restored with key: %s\n", alias, keyPath)
	return nil
}
