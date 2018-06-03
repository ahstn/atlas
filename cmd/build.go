package cmd

import (
	"github.com/ahstn/atlas/pb"
	"github.com/urfave/cli"
)

// Build defines the command for the cli and the logic to build the app
var Build = cli.Command{
	Name:    "build",
	Aliases: []string{"b"},
	Usage:   "execute the application build process",
	Action: func(c *cli.Context) error {
		defer pb.RunProgressBar("Building")
		return nil
	},
}
