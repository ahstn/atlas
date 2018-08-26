package cmd

import (
	"context"
	"path"

	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/ahstn/atlas/pkg/config"
	"github.com/ahstn/atlas/pkg/docker"
	"github.com/ahstn/atlas/pkg/validator"
	"github.com/urfave/cli"
)

// Docker defines the command for the cli and the logic to utilise Docker
// i.e atlas docker --tag atlas:1.0.0 --arg VERSION=0.1.0 --arg LANG=go'
var Docker = cli.Command{
	Name:      "docker",
	Aliases:   []string{"d"},
	Usage:     "build an application's Dockerfile",
	ArgsUsage: "[directory containing Dockerfile]",
	Action:    DockerAction,
	Subcommands: []cli.Command{
		{
			Name:      "run",
			Usage:     "run an application's Docker image",
			ArgsUsage: "[app name]",
			Action:    runAction,
		},
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "tag, t",
			Usage: "name and tag image in the `name:tag` format",
		},
		cli.StringSliceFlag{
			Name:  "arg, a",
			Usage: "build arguments in the `arg=value` format (space seperated)",
		},
		flag.Config,
		flag.Verbose,
	},
}

func runAction(c *cli.Context) error {
	artifact := config.DockerArtifact{
		Tag: "product-api:test",
	}
	ctx := context.Background()
	return docker.RunContainer(ctx, artifact)
}

// DockerAction handles building a container
func DockerAction(c *cli.Context) error {
	ctx := context.Background()

	p, err := validator.ValidateArguments(c.Args().First())
	if err != nil {
		panic(err)
	}

	err = validator.ValidateTag(c.String("tag"))
	if err != nil {
		panic(err)
	}

	err = validator.ValidateBuildArgs(c.StringSlice("arg"))
	if err != nil {
		panic(err)
	}

	artifact := config.DockerArtifact{
		Tag:        c.String("tag"),
		Args:       c.StringSlice("args"),
		Path:       p,
		Dockerfile: path.Join(p, "Dockerfile"),
	}
	return docker.ImageBuild(ctx, artifact)
}
