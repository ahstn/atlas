package flag

import "github.com/urfave/cli"

var Verbose = cli.BoolFlag{
	Name:  "verbose, V",
	Usage: "verbose logging rather than progress bars",
}
