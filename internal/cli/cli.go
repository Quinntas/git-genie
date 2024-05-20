package cli

import (
	"os"
)

type Args struct {
	Command command
	Args    []string
}

func NewArgs() *Args {
	if len(os.Args) == 1 {
		return nil
	}

	cmd := commandValidator(os.Args[1])

	return &Args{
		Command: cmd,
		Args:    os.Args[2:],
	}
}
