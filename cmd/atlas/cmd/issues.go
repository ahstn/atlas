package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ahstn/atlas/pkg/git"
	"github.com/ahstn/atlas/pkg/util"
	"github.com/urfave/cli"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

// Issues defines the command for the cli to open browser at Issues URL
var Issues = cli.Command{
	Name:    "issues",
	Aliases: []string{"i", "issues"},
	Usage:   "open JIRA/Github issue page for current Git project",
	Action:  IssuesAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Usage: "private JIRA base `URL`(without '/issues' or '<team>')",
		},
	},
}

// IssuesAction executes logic to determine URL for 'Issues' page
func IssuesAction(c *cli.Context) error {
	url, err := determineURL(c)
	if err != nil {
		panic(err)
	}

	out, err := git.Branch(os.Getenv("PWD"))
	if err != nil {
		panic(err)
	}

	branch := strings.TrimSpace(strings.ToLower(string(out)))
	if git.IsShortLivedBranch(branch) {
		branch = git.TrimBranchContext(branch)

		issueID := strings.SplitAfter(branch, "/")[1]
		url = fmt.Sprintf("%s/%s", url, issueID)
		emoji.Printf(":globe_with_meridians:Opening Repo Issue URL: %v", url)
	} else {
		emoji.Printf(":globe_with_meridians:Opening Feature Issue URL: %v", url)
	}

	return util.OpenBrowser(url)
}

func determineURL(c *cli.Context) (string, error) {
	var url string

	if c.IsSet("url") {
		url = c.String("url")
	} else {
		out, err := git.URL(os.Getenv("PWD"))
		if err != nil {
			return "", err
		}
		url, err = util.ProcessRepoURL(string(out))
		if err != nil {
			return "", nil
		}
	}

	return fmt.Sprintf("%s/issues", strings.TrimSpace(url)), nil
}
