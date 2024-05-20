package cmd

import "os/exec"

func Run(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(stdout), nil
}
