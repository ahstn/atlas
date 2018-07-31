package git

import (
	"os/exec"
	"strings"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

// BranchLogMessage establish branch type from given branch name
func BranchLogMessage(branch, url string) string {
	if strings.Contains(branch, "develop") || strings.Contains(branch, "master") {
		return emoji.Sprintf(":globe_with_meridians:Opening Repo Issue URL: %v", url)
	}
	return emoji.Sprintf(":globe_with_meridians:Opening Feature Issue URL: %v", url)
}

// Branch getter function for git branch
func Branch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	branch := string(out)
	//TODO remove the ToLower here, it would cause issues if issue name was CAPS e.g. DEV-118
	return strings.TrimSpace(strings.ToLower(branch)), nil
}
