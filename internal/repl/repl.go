package repl

import (
	"bufio"
	"fmt"
	"minishell/internal/parser"
	"os"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		dir, err := os.Getwd()
		if err != nil {
			return
		}
		fmt.Printf("%s $ ", dir)

		if !scanner.Scan() {
			fmt.Println("\nexit")
			break
		}

		line := scanner.Text()
		if line == "" {
			continue
		}

		cmd, err := parser.Parse(line)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		if cmd == nil {
			continue
		}

		if err := cmd.Execute(); err != nil {
			fmt.Println("error:", err)
		}
	}
}
