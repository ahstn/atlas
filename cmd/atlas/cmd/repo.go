package cmd

import (
	"fmt"
	"os"

	"github.com/ahstn/atlas/pkg/git"
	"github.com/ahstn/atlas/pkg/log"
	"github.com/ahstn/atlas/pkg/util"
	"github.com/urfave/cli"
)

// Repo defines the command for the cli to open a repo's URL
var Repo = cli.Command{
	Name:    "repo",
	Aliases: []string{"r"},
	Usage:   "open Git repo in browser",
	Action: func(c *cli.Context) error {
		logger := log.NewClient()
		return repo(c, logger, new(git.Client))
	},
}

func repo(c *cli.Context, logger log.Logger, git git.SourceController) error {
	out, err := git.URL(os.Getenv("PWD"))
	logger.CheckError(err)

	url, err := util.ProcessRepoURL(string(out))
	logger.CheckError(err)

	logger.Print(fmt.Sprintf(":globe_with_meridians:Opening Repo URL: %v \n", url))
	return util.OpenBrowser(url)
}
