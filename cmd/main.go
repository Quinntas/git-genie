package main

import (
	"github.com/quinntas/git-genie/internal/cli"
	"github.com/quinntas/git-genie/internal/gh"
	"github.com/quinntas/git-genie/internal/git"
	stringUtils "github.com/quinntas/git-genie/internal/utils/string"
)

func verifyPackages() {
	err := git.VerifyInstallation()
	if err != nil {
		panic(err)
	}

	err = gh.VerifyInstallation()
	if err != nil {
		panic(err)
	}
}

func main() {
	verifyPackages()

	args := cli.NewArgs()

	if args != nil {
		switch (*args).Command {
		case cli.COMMAND_COMMIT_ALL:
			err := git.CommitAll(stringUtils.ArrayToString((*args).Args))
			if err != nil {
				panic(err)
			}
		default:
			panic("unhandled default case")
		}
	}

	//p := tea.NewProgram(models.NewMainModel())
	//if _, err := p.Run(); err != nil {
	//	panic(err)
	//}
}
