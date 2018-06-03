package cmd

import (
	"github.com/ahstn/atlas/pb"
	"github.com/urfave/cli"
)

// Clean defines the command for the cli
var Clean = cli.Command{
	Name:    "clean",
	Aliases: []string{"c"},
	Usage:   "clean the applications build artifacts and dependencies",
	Action:  CleanAction,
	Subcommands: []cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "executes the application build process",
			Action: func(c *cli.Context) error {
				if err := CleanAction(c); err != nil {
					return err
				}
				return BuildAction(c)
			},
		},
	},
}

// CleanAction executes the logic to clean the application build environment
func CleanAction(c *cli.Context) error {
	defer pb.RunProgressBar("Cleaning")
	return nil
}
