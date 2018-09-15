package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ahstn/atlas/pkg/git"
	"github.com/ahstn/atlas/pkg/log"
	"github.com/ahstn/atlas/pkg/util"
	"github.com/urfave/cli"
)

// Issues defines the command for the cli to open browser at Issues URL
var Issues = cli.Command{
	Name:    "issues",
	Aliases: []string{"i"},
	Usage:   "open JIRA/Github issue tracker for current Git project",
	Action: func(c *cli.Context) error {
		logger := log.NewClient()
		return issues(c, logger, new(git.Client))
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Usage: "private JIRA base `URL` (without '/issues' or '<team>')",
		},
	},
}

// issues is the command action and allows dependency injection for testing
func issues(c *cli.Context, logger log.Logger, g git.SourceController) error {
	url, err := determineURL(c, g)
	logger.CheckError(err)

	out, err := g.Branch(os.Getenv("PWD"))
	logger.CheckError(err)

	branch := strings.TrimSpace(strings.ToLower(string(out)))
	if git.IsShortLivedBranch(branch) {
		branch = git.TrimBranchContext(branch)

		issueID := strings.SplitAfter(branch, "/")[1]
		url = fmt.Sprintf("%s/%s", url, issueID)
		logger.Printf(":globe_with_meridians:Opening Repo Issue URL: %v", url)
	} else {
		logger.Printf(":globe_with_meridians:Opening Feature Issue URL: %v", url)
	}

	return util.OpenBrowser(url)
}

func determineURL(c *cli.Context, g git.SourceController) (string, error) {
	var url string

	if c.IsSet("url") {
		url = c.String("url")
	} else {
		out, err := g.URL(os.Getenv("PWD"))
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
