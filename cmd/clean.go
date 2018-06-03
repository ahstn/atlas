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
}

// CleanAction executes the logic to clean the application build environment
// While also allowing commands to be chained
// i.e. "atlas clean build"
func CleanAction(c *cli.Context) error {
	cliArgs := c.Args()

	pb.RunProgressBar("Cleaning")

	for _, a := range cliArgs {
		if Build.HasName(a) {
			defer Build.Run(c)
		} else if Package.HasName(a) {
			defer Package.Run(c)
		}
	}
	return nil
}
