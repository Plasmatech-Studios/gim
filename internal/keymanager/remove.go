package keymanager

import (
	"bufio"
	"fmt"
	"gim/internal/config"
	"gim/internal/utils"
	"os"
	"path/filepath"
)

func RemoveKey(alias string, deleteFiles bool) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	keyPath, exists := cfg.Aliases[alias]
	if !exists {
		fmt.Printf("Alias '%s' not found in config. Checking for orphaned key...\n", alias)
		keyPath = filepath.Join(utils.GetSSHDir(), alias)

		if !utils.IsFileExists(keyPath) {
			return fmt.Errorf("key '%s' not found in config or filesystem", alias)
		}

		fmt.Printf("Alias '%s' was orphaned (not found in config).\n", alias)
	}

	// Always try to remove from config, no harm if it doesn't exist.
	delete(cfg.Aliases, alias)
	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to save updated config: %w", err)
	}
	if exists {
		fmt.Printf("Alias '%s' removed from config.\n", alias)
	}

	// If deleteFiles is false, skip file deletion.
	if deleteFiles {
		if err := deleteKeyFiles(keyPath); err != nil {
			fmt.Printf("Warning: Could not delete key files: %v\n", err)
		} else {
			fmt.Printf("Key files for '%s' deleted successfully.\n", alias)
		}
	}

	return nil
}


func deleteKeyFiles(keyPath string) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Are you sure you want to delete the key files for '%s'? (y/N): ", keyPath)
	response, _ := reader.ReadString('\n')

	if response != "y\n" && response != "Y\n" {
		return fmt.Errorf("deletion declined")
	}

	if err := os.Remove(keyPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete private key: %w", err)
	}

	pubKeyPath := keyPath + ".pub"
	if err := os.Remove(pubKeyPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete public key: %w", err)
	}

	return nil
}
