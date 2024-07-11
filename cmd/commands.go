package cmd

import (
	"errors"
)

var commands map[string]Command = map[string]Command{
	"add":      NewAddCmd(),
	"remove":   NewRemoveCmd(),
	"list":     NewList(),
	"copypath": NewCopyPathCmd(),
	"help":     NewHelpCmd(),
	"contains": NewContainsCmd(),
}

func GetCommand(args []string) (Command, error) {

	if cmd, exists := commands[args[1]]; exists {
		return cmd, nil
	}
	return nil, errors.New("command not found")
}
