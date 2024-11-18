package commands

import (
	"fmt"
	"gim/internal/keymanager"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Handler     func(args []string) error
}

var Commands = map[string]Command{
	"list": {
		Name:        "list",
		Description: "Lists all configured keys. Use -a to include orphaned keys.",
		Usage:       "gim list [-a]",
		Handler: func(args []string) error {
			showAll := len(args) > 0 && args[0] == "-a"
			return keymanager.ListKeys(showAll)
		},
	},
	"add": {
		Name:        "add",
		Description: "Adds a new SSH key with the given alias.",
		Usage:       "gim add <alias>",
		Handler: func(args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("missing alias")
			}
			return keymanager.AddKey(args[0])
		},
	},
	"use": {
		Name:        "use",
		Description: "Switches to the specified SSH key.",
		Usage:       "gim use <alias>",
		Handler: func(args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("missing alias")
			}
			return keymanager.UseKey(args[0])
		},
	},
	"remove": {
		Name:        "remove",
		Description: "Removes a key alias. Use -d to delete associated files.",
		Usage:       "gim remove [-d] <alias>",
		Handler: func(args []string) error {
			deleteFiles := false
			alias := ""

			if len(args) > 0 && args[0] == "-d" {
				deleteFiles = true
				if len(args) < 2 {
					return fmt.Errorf("missing alias")
				}
				alias = args[1]
			} else if len(args) > 0 {
				alias = args[0]
			} else {
				return fmt.Errorf("missing alias")
			}

			return keymanager.RemoveKey(alias, deleteFiles)
		},
	},
	"restore": {
		Name:        "restore",
		Description: "Restores an orphaned key alias.",
		Usage:       "gim restore <alias>",
		Handler: func(args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("missing alias")
			}
			return keymanager.RestoreKey(args[0])
		},
	},
	"rename": {
		Name:        "rename",
		Description: "Renames an existing alias.",
		Usage:       "gim rename <oldAlias> <newAlias>",
		Handler: func(args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("missing oldAlias or newAlias")
			}
			return keymanager.RenameAlias(args[0], args[1])
		},
	},
	"using": {
		Name:        "using",
		Description: "Displays the currently active SSH key. Use -c to copy the public key.",
		Usage:       "gim using [-c]",
		Handler: func(args []string) error {
			copyPublicKey := len(args) > 0 && args[0] == "-c"
			return keymanager.GetActiveKey(copyPublicKey)
		},
	},
}
