package cmd

import (
	"os/exec"

	"github.com/ahstn/atlas/util"
	"github.com/urfave/cli"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

// Repo defines the command for the cli to open a repo's URL
var Repo = cli.Command{
	Name:    "repo",
	Aliases: []string{"r"},
	Usage:   "open Git repo in browser",
	Action:  RepoAction,
}

// RepoAction executes the logic to open a repo's URL
func RepoAction(c *cli.Context) error {
	cmd := exec.Command("git", "ls-remote", "--get-url")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	url, err := util.ProcessRepoURL(string(out))
	if err != nil {
		panic(err)
	}

	emoji.Printf(":globe_with_meridians:Opening Repo URL: %v \n", url)
	util.OpenBrowser(url)
	if err != nil {
		panic(err)
	}
	return nil
}
