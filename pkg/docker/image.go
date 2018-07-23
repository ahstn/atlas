package docker

import (
	"context"
	"io"

	"github.com/ahstn/atlas/pkg/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	apiSocket  = "unix:///var/run/docker.sock"
	apiVersion = "v1.24"
)

var (
	apiHeaders = map[string]string{"User-Agent": "engine-api-cli-1.0"}
)

// ImageBuild takes a DockerArtifact and uses the Docker daemon for building
func ImageBuild(c context.Context) error {
	r, w := io.Pipe()
	go func() {
		err := util.CreateArchive(".", w)
		if err != nil {
			panic(err)
		}
		w.Close()
	}()

	opts := types.ImageBuildOptions{
		Tags:       []string{"atlas-dockerfile:test"},
		Dockerfile: "./Dockerfile",
		BuildArgs:  map[string]*string(nil),
	}

	cli, err := client.NewClient(apiSocket, apiVersion, nil, apiHeaders)
	res, err := cli.ImageBuild(c, r, opts)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	return PrintStream(res.Body)
}
