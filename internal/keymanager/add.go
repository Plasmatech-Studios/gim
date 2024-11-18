package keymanager

import (
	"fmt"
	"gim/internal/config"
	"gim/internal/utils"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/atotto/clipboard"
)

func AddKey(alias string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	sshDir := utils.GetSSHDir()
	privateKeyPath := filepath.Join(sshDir, alias)
	publicKeyPath := privateKeyPath + ".pub"

	// Generate private and public key
	cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-f", privateKeyPath, "-q", "-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate SSH key: %w", err)
	}

	// Verify public key was created
	if _, err := os.Stat(publicKeyPath); os.IsNotExist(err) {
		return fmt.Errorf("public key not created: %s", publicKeyPath)
	}

	pubKeyContent, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return fmt.Errorf("could not read public key: %w", err)
	}
	if err := clipboard.WriteAll(string(pubKeyContent)); err != nil {
		return fmt.Errorf("failed to copy public key to clipboard: %w", err)
	}

	// Add the new key to the config
	cfg.Aliases[alias] = privateKeyPath
	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	fmt.Printf("Key added for alias '%s':\n  Private: %s\n  Public:  %s\n", alias, privateKeyPath, publicKeyPath)
	fmt.Println("\nPublic key copied to clipboard.")
	
	return nil
}
