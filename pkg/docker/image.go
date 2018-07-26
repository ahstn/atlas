package docker

import (
	"context"
	"io"
	"strings"

	"github.com/ahstn/atlas/pkg/config"
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
// TODO: Path and Dockerfile must be relative (i.e. "./" and "./Dockerfile")
func ImageBuild(c context.Context, d config.DockerArtifact) error {
	r, w := io.Pipe()
	go func() {
		err := util.CreateArchive(d.Path, w)
		if err != nil {
			panic(err)
		}
		w.Close()
	}()

	args, err := ParseBuildArgsFromFlag(d.Args)
	if err != nil {
		panic(err)
	}

	opts := types.ImageBuildOptions{
		Tags:       []string{d.Tag},
		Dockerfile: d.Dockerfile,
		BuildArgs:  args,
	}

	cli, err := client.NewClient(apiSocket, apiVersion, nil, apiHeaders)
	res, err := cli.ImageBuild(c, r, opts)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	return PrintStream(res.Body)
}

// ParseBuildArgsFromFlag writes args from a flag string array to a map
// Example input: []string{"VERSION=1.8.0", "LANG=Go", "BUILDER=Docker"}
func ParseBuildArgsFromFlag(s []string) (map[string]*string, error) {
	if len(s) == 0 {
		return map[string]*string(nil), nil
	}

	m := make(map[string]*string)
	for _, arg := range s {
		a := strings.Split(strings.TrimSpace(arg), "=")
		m[a[0]] = &a[1]
	}

	return m, nil
}
