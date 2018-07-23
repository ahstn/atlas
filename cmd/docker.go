package cmd

import (
	"context"

	"github.com/ahstn/atlas/pkg/docker"
	"github.com/urfave/cli"
)

// Docker defines the command for the cli and the logic to utilise Docker
var Docker = cli.Command{
	Name:    "docker",
	Aliases: []string{"d"},
	Usage:   "execute the application build process",
	Action:  DockerAction,
}

// DockerAction handles building a container
func DockerAction(c *cli.Context) error {
	ctx := context.Background()

	return docker.ImageBuild(ctx)
}
