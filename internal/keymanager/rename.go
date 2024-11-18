package keymanager

import (
	"fmt"
	"gim/internal/config"
)

func RenameAlias(oldAlias, newAlias string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	keyPath, exists := cfg.Aliases[oldAlias]
	if !exists {
		return fmt.Errorf("alias '%s' not found in config", oldAlias)
	}

	// Check if the new alias already exists
	if _, exists := cfg.Aliases[newAlias]; exists {
		return fmt.Errorf("alias '%s' already exists in config", newAlias)
	}

	cfg.Aliases[newAlias] = keyPath
	delete(cfg.Aliases, oldAlias)

	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to save updated config: %w", err)
	}

	fmt.Printf("Alias '%s' renamed to '%s'.\n", oldAlias, newAlias)
	return nil
}
