package main

import (
	"github.com/quinntas/git-genie/internal/cli"
	"github.com/quinntas/git-genie/internal/gh"
	"github.com/quinntas/git-genie/internal/git"
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

	err := cli.HandleCli()
	if err != nil {
		panic(err)
	}

	//p := tea.NewProgram(models.NewMainModel())
	//if _, err := p.Run(); err != nil {
	//	panic(err)
	//}
}
