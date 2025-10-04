package parser

import (
	"fmt"
	"minishell/internal/commands"
	"strconv"
	"strings"
)

func Parse(line string) (commands.Command, error) {
	tokens := strings.Fields(line)
	if len(tokens) == 0 {
		return nil, nil
	}

	switch tokens[0] {
	case "pwd":
		return &commands.PwdCommand{}, nil
	case "cd":
		path := ""
		if len(tokens) > 1 {
			path = tokens[1]
		}
		return &commands.CdCommand{Path: path}, nil
	case "echo":
		return &commands.EchoCommand{Args: tokens[1:]}, nil
	case "kill":
		if len(tokens) < 2 {
			return nil, fmt.Errorf("kill: missing pid")
		}
		pid, err := strconv.Atoi(tokens[1])
		if err != nil {
			return nil, fmt.Errorf("kill: invalid pid")
		}
		return &commands.KillCommand{Pid: pid}, nil
	case "ps":
		return &commands.PsCommand{}, nil
	case "exit":
		return &commands.ExitCommand{}, nil
	default:
		return &commands.ExternalCommand{Name: tokens[0], Args: tokens[1:]}, nil
	}
}
