package cmd

import (
	"os/exec"
	"strings"

	"github.com/ahstn/atlas/pkg/git"
	"github.com/ahstn/atlas/pkg/util"
	"github.com/urfave/cli"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

/*
This whole file needs reworking.
Initial commit to verify refactoring hasn't broken anything
+ initial test of feature.
*/

// Issues defines the command for the cli to open browser at Issues URL
var Issues = cli.Command{
	Name:    "issues",
	Aliases: []string{"i", "issues"},
	Usage:   "Open Jira/Github issue page for current Git project",
	Action:  IssuesAction,
}

// IssuesAction executes logic to determine URL for 'Issues' page
func IssuesAction(c *cli.Context) error {

	//Consider pulling these into pkg/
	//Lines 31-35 -> git/getUrl.go
	cmd := exec.Command("git", "ls-remote", "--get-url")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	url, err := util.ProcessRepoURL(string(out))
	if err != nil {
		panic(err)
	}

	url = strings.Replace(url, ".git", "/issues", 1)

	//Lines 46-50 -> git/getBranch.go
	cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err = cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	branch := string(out[:])

	branch = git.DetermineBranchType(string(branch))
	if strings.Contains(branch, "feature") {
		emoji.Printf(":globe_with_meridians:Opening Feature Issue URL: %v \n", url)
	} else {
		emoji.Printf(":globe_with_meridians:Opening Repo Issue URL: %v \n", url)
	}
	util.OpenBrowser(url)

	if err != nil {
		panic(err)
	}
	return nil
}
