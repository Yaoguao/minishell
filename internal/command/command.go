package command

import "io"

type Command interface {
	Run(input io.Reader, output io.Writer) error
	IsBuiltin() bool
}
