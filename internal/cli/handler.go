package cli

import (
	"errors"
	"github.com/quinntas/git-genie/internal/git"
	stringUtils "github.com/quinntas/git-genie/internal/utils/string"
)

func HandleCli() error {
	args := NewArgs()

	if args != nil {
		switch (*args).Command {
		case COMMAND_COMMIT_ALL:
			err := git.CommitAll(stringUtils.ArrayToString((*args).Args))
			if err != nil {
				return err
			}

		case COMMAND_PUSH:
			err := git.Push()
			if err != nil {
				return err
			}

		default:
			return errors.New("command not handled")
		}
	}

	return nil
}
