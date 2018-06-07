package flag

import "github.com/urfave/cli"

var SkipTests = cli.BoolTFlag{
	Name:  "s, skipTests",
	Usage: "skip tests",
}
