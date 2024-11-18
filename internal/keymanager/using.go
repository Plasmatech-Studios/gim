package keymanager

import (
	"fmt"
	"gim/internal/config"
	"os"
	"os/exec"
	"strings"

	"github.com/atotto/clipboard"
)

func GetActiveKey(copyPublicKey bool) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	cmd := exec.Command("ssh-add", "-l")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list active keys: %w", err)
	}

	activeKeys := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(activeKeys) == 0 || (len(activeKeys) == 1 && strings.Contains(activeKeys[0], "no identities")) {
		fmt.Println("No active SSH keys in the agent.")
		return nil
	}

	fmt.Println("\nActive SSH Keys:")
	for _, key := range activeKeys {
		parts := strings.Fields(key)
		if len(parts) < 2 {
			continue
		}

		fingerprint := parts[1]
		alias := findAliasByFingerprint(cfg, fingerprint)
		if alias != "" {
			keyPath := cfg.Aliases[alias]
			fmt.Printf("  %-10s -> %s\n", alias, keyPath)

			if copyPublicKey {
				pubKeyPath := keyPath + ".pub"
				pubKeyContent, err := os.ReadFile(pubKeyPath)
				if err != nil {
					return fmt.Errorf("failed to read public key: %w", err)
				}

				if err := clipboard.WriteAll(string(pubKeyContent)); err != nil {
					return fmt.Errorf("failed to copy public key to clipboard: %w", err)
				}

				fmt.Printf("Public key for '%s' copied to clipboard.\n", alias)
			}
		} else {
			fmt.Printf("  (unknown)   -> %s\n", fingerprint)
		}
	}

	return nil
}

func findAliasByFingerprint(cfg *config.Config, fingerprint string) string {
	for alias, path := range cfg.Aliases {
		if keyMatchesFingerprint(path, fingerprint) {
			return alias
		}
	}
	return ""
}

func keyMatchesFingerprint(path, fingerprint string) bool {
	cmd := exec.Command("ssh-keygen", "-lf", path)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(output), fingerprint)
}
