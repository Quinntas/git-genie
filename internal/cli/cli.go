package cli

import (
	"flag"
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

	cmdPtr := flag.String("cmd", "", "Git command")

	flag.Parse()

	cmd := commandValidator(*cmdPtr)

	return &Args{
		Command: cmd,
		Args:    flag.Args(),
	}
}
