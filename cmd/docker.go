package cmd

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/urfave/cli"
)

// Docker defines the command for the cli and the logic to utilise Docker
var Docker = cli.Command{
	Name:    "docker",
	Aliases: []string{"d"},
	Usage:   "execute the application build process",
	Action:  DockerAction,
}

const (
	apiSocket  = "unix:///var/run/docker.sock"
	apiVersion = "v1.24"
)

var (
	apiHeaders = map[string]string{"User-Agent": "engine-api-cli-1.0"}
)

// DockerAction handles building a container
func DockerAction(c *cli.Context) error {
	// Create types.ImageBuildOptions from config and flags
	// Compress directory including Dockerfile into tar
	// Send tar as Docker Context using ImageBuild()
	ctx := context.Background()
	reader, writer := io.Pipe()

	go func() {
		err := CreateArchive(".", writer)
		if err != nil {
			panic(err)
		}
		writer.Close()
	}()

	opts := types.ImageBuildOptions{
		Tags:       []string{"atlas-dockerfile:test"},
		Dockerfile: "./Dockerfile",
		BuildArgs:  map[string]*string(nil),
	}

	cli, err := client.NewClient(apiSocket, apiVersion, nil, apiHeaders)
	res, err := cli.ImageBuild(ctx, reader, opts)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyBytes))

	return nil
}

// CreateArchive creates a .tar archive to be passed into Docker's ImageBuild()
func CreateArchive(p string, w io.Writer) error {
	tw := tar.NewWriter(w)
	defer tw.Close()

	return filepath.Walk(p, func(file string, fi os.FileInfo, err error) error {
		if err != nil || !fi.Mode().IsRegular() {
			return err
		}

		fmt.Println(file)
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}
		header.Name = file
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		return nil
	})
}
