package git

import (
	"fmt"
	"github.com/quinntas/git-genie/internal/utils/cmd"
)

func VerifyInstallation() error {
	_, err := cmd.Run("git", "--version")
	if err != nil {
		return fmt.Errorf("git is not installed")
	}
	return nil
}

func Add(path string) error {
	_, err := cmd.Run("git", "add", path)
	if err != nil {
		return fmt.Errorf("failed to add %s", path)
	}
	return nil
}

func Commit(message string) error {
	out, err := cmd.Run("git", "commit", "-m", message)
	if err != nil {
		return fmt.Errorf("failed to commit")
	}
	fmt.Println(out)
	return nil
}

func CommitAll(message string) error {
	err := Add(".")
	if err != nil {
		return err
	}
	err = Commit(message)
	if err != nil {
		return err
	}
	return nil
}

func Push() error {
	out, err := cmd.Run("git", "push")
	if err != nil {
		return fmt.Errorf("failed to push")
	}
	fmt.Println(out)
	return nil
}

func Clone(url string) error {
	_, err := cmd.Run("git", "clone", url)
	if err != nil {
		return fmt.Errorf("failed to clone - %s", url)
	}
	return nil
}
