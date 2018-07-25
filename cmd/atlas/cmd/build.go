package cmd

import (
	"os"

	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/ahstn/atlas/pkg/builder"
	"github.com/apex/log"
	"github.com/urfave/cli"
)

// Build defines the command for the cli and the logic to build the app
var Build = cli.Command{
	Name:    "build",
	Aliases: []string{"b"},
	Usage:   "execute the application build process",
	Action:  BuildAction,
	Flags: []cli.Flag{
		flag.Clean,
		flag.SkipTests,
		flag.Verbose,
	},
}

// BuildAction executes the logic to clean the application build environment
// While also allowing commands to be chained
// i.e. "atlas clean build"
func BuildAction(c *cli.Context) error {
	var mvn builder.Maven

	if c.Bool("clean") {
		mvn.Clean()
	}

	mvn.Build()

	if err := mvn.Run(c.Bool("verbose")); err != nil {
		log.Info("Error:" + err.Error())
		os.Exit(1)
	}

	return nil
}
