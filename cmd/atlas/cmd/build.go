package cmd

import (
	"os"

	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/ahstn/atlas/pkg/builder"
	"github.com/urfave/cli"
)

// Build defines the command for the cli and the logic to build the app
var Build = cli.Command{
	Name:    "build",
	Aliases: []string{"b"},
	Usage:   "execute the application build process",
	Action: func(c *cli.Context) error {
		mvn := builder.NewClient(os.Getenv("PWD"), nil, nil, nil)
		return build(c, mvn)
	},
	Flags: []cli.Flag{
		flag.Clean,
		flag.SkipTests,
		flag.Verbose,
	},
}

// build executes the logic to clean the application build environment
// While also allowing commands to be chained
// i.e. "atlas clean build"
func build(c *cli.Context, b builder.Builder) error {
	goals := []string{"install"}

	if c.IsSet("clean") {
		goals = append([]string{"clean"}, goals...)
	}

	b.ModifyArgs(goals)
	return b.Run(c.Bool("verbose"))
}
