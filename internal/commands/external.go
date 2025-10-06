package commands

import (
	"io"
	"os"
	"os/exec"
)

type ExternalCommand struct {
	Name string
	Args []string
}

func (e *ExternalCommand) Run(input io.Reader, output io.Writer) error {
	cmd := exec.Command(e.Name, e.Args...)
	cmd.Stdin = input
	cmd.Stdout = output
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (e *ExternalCommand) IsBuiltin() bool { return false }
