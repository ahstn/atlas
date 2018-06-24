package flag

import "github.com/urfave/cli"

var Verbose = cli.BoolFlag{
	Name:  "V, verbose",
	Usage: "verbose logging rather than progress bars",
}
