package cmd

import (
	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/urfave/cli"
)

// Git defines the command for the preforming git actions against services
// i.e atlas git checkout master -e app1
var Git = cli.Command{
	Name:    "git",
	Aliases: []string{"g"},
	Usage:   "preform Git actions against service(s)",
	Action:  DockerAction,
	Subcommands: []cli.Command{
		{
			Name:      "clone",
			Usage:     "clone the services' repo(s) defined in config",
			ArgsUsage: "[single service]",
		},
		{
			Name:      "branch",
			Usage:     "create a branch in the services' repo(s) defined in config",
			ArgsUsage: "[branch name]",
		},
		{
			Name:      "checkout",
			Usage:     "checkout a branch in services' repo(s) defined in config",
			ArgsUsage: "[branch]",
		},
	},
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:  "exclude, e",
			Usage: "exclude certain services defined in config from the command",
		},
		flag.Config,
	},
}

func GitAction() error {
	return nil
}
