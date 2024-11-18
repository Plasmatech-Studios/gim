package keymanager

import (
	"fmt"
	"gim/internal/config"
	"os/exec"
)

func UseKey(alias string) error {
    cfg, err := config.LoadConfig()
    if err != nil {
        return fmt.Errorf("failed to load config: %w", err)
    }

    keyPath, exists := cfg.Aliases[alias]
    if !exists {
        return fmt.Errorf("alias '%s' not found in config", alias)
    }

    if err := removeAllKeys(); err != nil {
        return fmt.Errorf("failed to remove existing keys: %w", err)
    }

    if err := addKeyToAgent(keyPath); err != nil {
        return fmt.Errorf("failed to add key '%s' to agent: %w", alias, err)
    }

    // Update the config to mark this key as "using"
    cfg.Using = alias
    if err := config.SaveConfig(cfg); err != nil {
        return fmt.Errorf("failed to update config: %w", err)
    }

    fmt.Printf("Successfully switched to key: %s (%s)\n", alias, keyPath)
    return nil
}

func removeAllKeys() error {
	cmd := exec.Command("ssh-add", "-D")
	cmd.Stdout = nil // Suppress output
	cmd.Stderr = nil
	return cmd.Run()
}

func addKeyToAgent(keyPath string) error {
	cmd := exec.Command("ssh-add", keyPath)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
