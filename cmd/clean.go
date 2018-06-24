package cmd

import (
	"os"

	"github.com/ahstn/atlas/builder"
	"github.com/ahstn/atlas/flag"
	"github.com/apex/log"
	"github.com/urfave/cli"
)

// Clean defines the command for the cli
var Clean = cli.Command{
	Name:    "clean",
	Aliases: []string{"c"},
	Usage:   "clean the applications build artifacts and dependencies",
	Action:  CleanAction,
	Flags: []cli.Flag{
		flag.SkipTests,
		flag.Verbose,
	},
}

// CleanAction executes the logic to clean the application build environment
// While also allowing commands to be chained
// i.e. "atlas clean build"
func CleanAction(c *cli.Context) error {
	cliArgs := c.Args()
	var mvn builder.Maven
	mvn.Clean()

	for _, a := range cliArgs {
		if Build.HasName(a) {
			defer Build.Run(c)
			mvn.Build()
		} else if Package.HasName(a) {
			defer Package.Run(c)
			mvn.Package()
		}
	}

	if err := mvn.Run(); err != nil {
		log.Info("Error:" + err.Error())
		os.Exit(1)
	}

	return nil
}
