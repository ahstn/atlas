package git

import (
	"strings"
)

// DetermineBranchType establish branch type given output from 'git branch'
func DetermineBranchType(branch string) string {
	branch = strings.ToLower(branch)

	if strings.Contains(branch, "develop") || strings.Contains(branch, "master") {
		return "develop"
	}

	return "feature"

}
