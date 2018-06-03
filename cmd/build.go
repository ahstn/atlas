package cmd

import (
	"github.com/ahstn/atlas/pb"
	"github.com/urfave/cli"
)

// Build defines the command for the cli
var Build = cli.Command{
	Name:    "build",
	Aliases: []string{"b"},
	Usage:   "execute the application build process",
	Action:  BuildAction,
}

// BuildAction executes the logic to build the application
func BuildAction(c *cli.Context) error {
	defer pb.RunProgressBar("Building")
	return nil
}
