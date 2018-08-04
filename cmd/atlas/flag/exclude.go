package flag

import "github.com/urfave/cli"

var Exclude = cli.StringSliceFlag{
	Name:  "exclude, e",
	Usage: "exclude certain services defined in config from the command",
}
