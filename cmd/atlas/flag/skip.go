package flag

import "github.com/urfave/cli"

var SkipTests = cli.BoolTFlag{
	Name:  "skipTests, s",
	Usage: "skip tests",
}
