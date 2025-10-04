package commands

import (
	"os"
	"os/exec"
)

type ExternalCommand struct {
	Name string
	Args []string
}

func (c *ExternalCommand) Execute() error {
	cmd := exec.Command(c.Name, c.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
