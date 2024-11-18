package keymanager

import (
	"fmt"
	"gim/internal/config"
	"gim/internal/utils"
	"os"
	"path/filepath"
)

func ListKeys(showAll bool) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	fmt.Println("\nConfigured SSH Keys:")
	if len(cfg.Aliases) == 0 {
		fmt.Println("  (None)")
	} else {
		for alias, path := range cfg.Aliases {
			fmt.Printf("  %-10s -> %s\n", alias, path)
		}
	}

	if showAll {
		fmt.Println("\nOrphaned SSH Keys (present in ~/.ssh but not in config):")
		sshDir := utils.GetSSHDir()
		files, err := os.ReadDir(sshDir)
		if err != nil {
			return fmt.Errorf("could not read .ssh directory: %w", err)
		}

		orphanFound := false
		for _, file := range files {
			fullPath := filepath.Join(sshDir, file.Name())
			if utils.IsPrivateKey(file.Name()) && !isInConfig(cfg, fullPath) {
				fmt.Printf("  %s\n", fullPath)
				orphanFound = true
			}
		}

		if !orphanFound {
			fmt.Println("  (None)")
		}
	}

	return nil
}

func isInConfig(cfg *config.Config, keyPath string) bool {
	for _, path := range cfg.Aliases {
		if path == keyPath {
			return true
		}
	}
	return false
}
