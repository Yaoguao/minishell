package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func render() {
	for {
		fmt.Print("> ")

		os.
			line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		tokens := strings.Fields(line)

		if isBuiltin(tokens[0]) {
			runBuiltin(tokens)
		} else {
			runExternal(tokens)
		}
	}
}

func main() {
	cmd := exec.Command("ls", "-l")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
