package git

import (
	"os/exec"
	"strings"
)

// URL getter function for git URL
func URL() (string, error) {
	cmd := exec.Command("git", "ls-remote", "--get-url")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	url := string(out)

	return strings.TrimSpace(url), nil
}
