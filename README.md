# miniShell

A Go mini-shell that supports built-in and external commands, pipelines, and environment variables.
The project is intended for educational purposes and to demonstrate the principles of shell's work.

## Functional

- **Built-in command:**
    - `cd [dir]` — change current dir.
    - `exit` — exit from shell.
    - `echo [args...]` — input text `$VAR`.
- **External command:** through `exec.Command`.
- **Pipeline:** connection command through`|`:
  ```bash
  cat file.txt | grep "pattern" | wc -l
  
## Build and start

```shell
git clone https://github.com/Yaoguao/minishell.git
cd minishell
go build -o minishell
./minishell
```
## Development recommendations

- Add support for the built-in export and unset commands.

- Support for redirects (>, <) and background tasks (&).

- Implementation of command history and alias.

- Adding tests for teams and pipelines.