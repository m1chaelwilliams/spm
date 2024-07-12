package cmd

import (
	"errors"
	"strings"
)

var commands map[string]Command = map[string]Command{
	"add":      NewAddCmd(),
	"remove":   NewRemoveCmd(),
	"list":     NewList(),
	"copypath": NewCopyPathCmd(),
	"help":     NewHelpCmd(),
	"contains": NewContainsCmd(),
	"edit":     NewEditCmd(),
	"lookup":   NewLookupCmd(),
}

func GetCommand(args []string) (Command, error) {

	arg := strings.Trim(args[1], "-")

	if cmd, exists := commands[arg]; exists {
		return cmd, nil
	}
	return nil, errors.New("command not found")
}
