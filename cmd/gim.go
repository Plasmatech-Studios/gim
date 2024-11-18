package main

import (
	"fmt"
	"gim/internal/commands"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	cmdName := os.Args[1]
	cmd, exists := commands.Commands[cmdName]
	if !exists {
		fmt.Printf("Unknown command: %s\n", cmdName)
		showHelp()
		return
	}

	if err := cmd.Handler(os.Args[2:]); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func showHelp() {
	fmt.Println("Usage: gim <command> [options]")
	fmt.Println("\nAvailable commands:")
	for _, cmd := range commands.Commands {
		fmt.Printf("  %-10s %s\n", cmd.Name, cmd.Description)
		fmt.Printf("    Usage: %s\n", cmd.Usage)
	}
}
