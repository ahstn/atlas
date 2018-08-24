package git

import (
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

// IsShortLivedBranch is syntax sugar for determining branch type (Git Flow)
func IsShortLivedBranch(b string) bool {
	return !strings.Contains(b, "develop") && !strings.Contains(b, "master")
}

// TrimBranchContext removes any extra context/info from the branch data
// Example: feature/TEAM-123-implement-feature -> feature/TEAM-123
func TrimBranchContext(b string) string {
	if strings.Count(b, "-") > 1 {
		s := strings.SplitAfter(b, "-")[:2]
		b = strings.Join(s, "")
		b = strings.TrimSuffix(b, "-")
	}

	return b
}
