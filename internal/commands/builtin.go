package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type PwdCommand struct{}

func (c *PwdCommand) Execute() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}

type CdCommand struct {
	Path string
}

func (c *CdCommand) Execute() error {
	if c.Path == "" || c.Path == "~" {
		home, _ := os.UserHomeDir()
		return os.Chdir(home)
	}
	return os.Chdir(c.Path)
}

type EchoCommand struct {
	Args []string
}

func (c *EchoCommand) Execute() error {
	fmt.Println(strings.Join(c.Args, " "))
	return nil
}

type KillCommand struct {
	Pid int
}

func (c *KillCommand) Execute() error {
	return syscall.Kill(c.Pid, syscall.SIGKILL)
}

type PsCommand struct{}

func (c *PsCommand) Execute() error {
	cmd := exec.Command("ps", "-e", "-o", "pid,comm")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type ExitCommand struct{}

func (c *ExitCommand) Execute() error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil
}
