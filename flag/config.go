package flag

import "github.com/urfave/cli"

// Config handles passing a file's path
// TODO: Validation, i.e. file ends in .yaml
var Config = cli.StringFlag{
	Name:  "c, config",
	Value: "atlas.yaml",
	Usage: "name of config file in ~/.config/atlas",
}
