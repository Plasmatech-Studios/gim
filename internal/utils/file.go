package utils

import (
	"os"
	"path/filepath"
)

// GetSSHDir retrieves the user's SSH directory.
func GetSSHDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Could not determine home directory")
	}
	return filepath.Join(home, ".ssh")
}

// IsPrivateKey checks if a file is a valid SSH private key.
func IsPrivateKey(filename string) bool {
	nonKeyFiles := map[string]bool{
		"known_hosts":     true,
		"config":          true,
		"authorized_keys": true,
	}
	if _, excluded := nonKeyFiles[filename]; excluded {
		return false
	}

	// Private keys typically have no extension
	return filepath.Ext(filename) == ""
}

// IsFileExists checks if a file exists.
func IsFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}