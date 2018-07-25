package cmd

import (
	"context"

	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/ahstn/atlas/pkg/docker"
	"github.com/urfave/cli"
)

// Docker defines the command for the cli and the logic to utilise Docker
// i.e docker build --tag atlas:1.0.0 --args VERSION=0.1.0, LANG=go
var Docker = cli.Command{
	Name:      "docker",
	Aliases:   []string{"d"},
	Usage:     "build an application's Dockerfile",
	ArgsUsage: "[directory containing Dockerfile]",
	Action:    DockerAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "tag, t",
			Usage: "name and tag image in the `name:tag` format",
		},
		cli.StringSliceFlag{
			Name:  "args, a",
			Usage: "build arguments in the `arg:value` format (comma seperated)",
		},
		flag.Config,
		flag.Verbose,
	},
}

// DockerAction handles building a container
func DockerAction(c *cli.Context) error {
	ctx := context.Background()

	return docker.ImageBuild(ctx)
}
