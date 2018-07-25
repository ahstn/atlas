package flag

import "github.com/urfave/cli"

var Clean = cli.BoolFlag{
	Name:  "clean, c",
	Usage: "clean artifacts before building",
}
