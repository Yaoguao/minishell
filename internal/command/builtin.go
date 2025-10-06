package command

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type BuiltinCommand struct {
	Name string
	Args []string
}

func (b *BuiltinCommand) Run(input io.Reader, output io.Writer) error {
	switch b.Name {
	case "cd":
		if len(b.Args) == 0 {
			return nil
		}
		return os.Chdir(b.Args[0])
	case "exit":
		os.Exit(0)
	case "echo":
		for i, arg := range b.Args {
			if strings.HasPrefix(arg, "$") && len(arg) > 1 {
				val := os.Getenv(arg[1:])
				fmt.Fprint(output, val)
			} else {
				fmt.Fprint(output, arg)
			}
			if i < len(b.Args)-1 {
				fmt.Fprint(output, " ")
			}
		}
		fmt.Fprintln(output)
		return nil
	default:
		return fmt.Errorf("unknown builtin: %s", b.Name)
	}
	return nil
}

func (b *BuiltinCommand) IsBuiltin() bool { return true }
