package parser

import (
	"io"
	"minishell/internal/commands"
	"strings"
)

type сommand interface {
	Run(input io.Reader, output io.Writer) error
	IsBuiltin() bool
}

func ParseCommand(cmdStr string) сommand {
	fields := strings.Fields(cmdStr)
	if len(fields) == 0 {
		return nil
	}

	name := fields[0]
	args := fields[1:]

	switch name {
	case "cd", "exit":
		return &commands.BuiltinCommand{Name: name, Args: args}
	default:
		return &commands.ExternalCommand{Name: name, Args: args}
	}
}
