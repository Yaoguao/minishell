package commands

import (
	"fmt"
	"io"
	"os"
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
	default:
		return fmt.Errorf("unknown builtin: %s", b.Name)
	}
	return nil
}

func (b *BuiltinCommand) IsBuiltin() bool { return true }
