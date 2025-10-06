package main

import (
	"bufio"
	"fmt"
	"io"
	"minishell/internal/command"
	"minishell/internal/parser"
	"os"
	"strings"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("minishell> ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		var cmds []command.Command
		for _, part := range parts {
			cmd := parser.ParseCommand(strings.TrimSpace(part))
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

		RunPipeline(cmds)
	}
}

func RunPipeline(cmds []command.Command) error {
	n := len(cmds)
	if n == 0 {
		return nil
	}

	pipes := make([]*io.PipeReader, n-1)
	pipeWriters := make([]*io.PipeWriter, n-1)
	for i := 0; i < n-1; i++ {
		r, w := io.Pipe()
		pipes[i] = r
		pipeWriters[i] = w
	}

	var wg sync.WaitGroup
	wg.Add(n)

	for i, cmd := range cmds {
		var input io.Reader
		var output io.Writer

		if i == 0 {
			input = os.Stdin
		} else {
			input = pipes[i-1]
		}

		if i == n-1 {
			output = os.Stdout
		} else {
			output = pipeWriters[i]
		}

		go func(cmd command.Command, input io.Reader, output io.Writer) {
			defer wg.Done()
			defer func() {
				if w, ok := output.(*io.PipeWriter); ok {
					w.Close()
				}
			}()

			if err := cmd.Run(input, output); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
			}
		}(cmd, input, output)
	}

	wg.Wait()
	return nil
}
