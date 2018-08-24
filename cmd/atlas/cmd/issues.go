package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ahstn/atlas/pkg/git"
	"github.com/ahstn/atlas/pkg/util"
	"github.com/urfave/cli"
)

// Issues defines the command for the cli to open browser at Issues URL
var Issues = cli.Command{
	Name:    "issues",
	Aliases: []string{"i", "issues"},
	Usage:   "Open Jira/Github issue page for current Git project",
	Action:  IssuesAction,
}

// IssuesAction executes logic to determine URL for 'Issues' page
func IssuesAction(c *cli.Context) error {
	out, err := git.URL(os.Getenv("PWD"))
	url, err := util.ProcessRepoURL(string(out))
	if err != nil {
		panic(err)
	}

	out, err = git.Branch(os.Getenv("PWD"))
	if err != nil {
		panic(err)
	}
	branch := strings.TrimSpace(strings.ToLower(string(out)))

	url = fmt.Sprintf("%s/issues", url)
	if git.IsShortLivedBranch(branch) {
		branch = git.TrimBranchContext(branch)

		branch = strings.SplitAfter(branch, "/")[1]
		url = fmt.Sprintf("%s/%s", url, branch)
	}

	fmt.Println(git.BranchLogMessage(branch, url))
	util.OpenBrowser(url)

	return nil
}
