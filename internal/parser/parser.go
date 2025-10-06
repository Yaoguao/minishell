package parser

import (
	"minishell/internal/command"
	"strings"
)

func ParseCommand(cmdStr string) command.Command {
	fields := strings.Fields(cmdStr)
	if len(fields) == 0 {
		return nil
	}

	name := fields[0]
	args := fields[1:]

	switch name {
	case "cd", "exit", "echo":
		return &command.BuiltinCommand{Name: name, Args: args}
	default:
		return &command.ExternalCommand{Name: name, Args: args}
	}
}
