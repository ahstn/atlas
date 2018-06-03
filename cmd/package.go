package cmd

import (
	"github.com/ahstn/atlas/pb"
	"github.com/urfave/cli"
)

// Package defines the command for the cli and the logic to package the app
var Package = cli.Command{
	Name:    "package",
	Aliases: []string{"p", "pkg"},
	Usage:   "package the application",
	Action: func(c *cli.Context) error {
		defer pb.RunProgressBar("Packaging")
		return nil
	},
}
