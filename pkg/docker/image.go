package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
func ImageBuild(c context.Context, d config.DockerArtifact) error {
	r, w := io.Pipe()
	err := os.Chdir(d.Path) // Must be in app dir when building
	if err != nil {
		return err
	}

	path, err := AbsToRelPath(d.Path)
	if err != nil {
		fmt.Println("[ImageBuild] Can't get rel")
		return err
	}

	dockerfile, err := AbsToRelPath(d.Dockerfile)
	if err != nil {
		return err
	}

	go func() {
		err := util.CreateArchive(path, w)
		if err != nil {
			panic(err)
		}
		w.Close()
	}()

	args, err := ParseBuildArgsFromFlag(d.Args)
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Tags:       []string{d.Tag},
		Dockerfile: dockerfile,
		BuildArgs:  args,
	}

	cli, err := client.NewClient(apiSocket, apiVersion, nil, apiHeaders)
	if err != nil {
		return err
	}

	res, err := cli.ImageBuild(c, r, opts)
	if err != nil {
		return err
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

// AbsToRelPath is a helper method for converting absolute paths to paths
// relative to the present working directory
// This is due to Docker requiring relative paths when building
func AbsToRelPath(p string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Rel(wd, p)
}
