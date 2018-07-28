package cmd

import (
	"fmt"
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
	url, err := git.URL()

	url, err = util.ProcessRepoURL(string(url))
	if err != nil {
		panic(err)
	}

	branch, err := git.Branch()
	if err != nil {
		panic(err)
	}

	url = strings.Replace(url, ".git", "/issues/", 1)
	if !strings.Contains(branch, "develop") && !strings.Contains(branch, "master") {
		branch = strings.SplitAfter(branch, "/")[1]
		url = url + branch
	}

	fmt.Println(git.BranchLogMessage(branch, url))
	util.OpenBrowser(url)

	return nil
}
