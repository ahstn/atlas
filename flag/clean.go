package flag

import "github.com/urfave/cli"

var Clean = cli.BoolFlag{
	Name:  "c, clean",
	Usage: "clean artifacts before building",
}
