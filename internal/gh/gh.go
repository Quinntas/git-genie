package gh

import (
	"fmt"
	"github.com/quinntas/git-genie/internal/utils/cmd"
)

func VerifyInstallation() error {
	_, err := cmd.Run("gh", "--version")
	if err != nil {
		return fmt.Errorf("gh is not installed")
	}
	return nil
}

func CreateRepo(name string) error {
	_, err := cmd.Run("gh", "repo", "create", name)
	if err != nil {
		return fmt.Errorf("failed to create repo")
	}
	return nil
}
