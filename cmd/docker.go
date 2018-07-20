package cmd

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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
	fmt.Println("Docker")
	// Create types.ImageBuildOptions from config and flags
	// Compress directory including Dockerfile into tar
	// Send tar as Docker Context using ImageBuild()
	err := CreateArchive(".")
	if err != nil {
		panic(err)
	}
	return nil
}

// CreateArchive creates a .tar archive to be passed into Docker's ImageBuild()
func CreateArchive(p string) error {
	buffer := bytes.NewBuffer(nil)
	tw := tar.NewWriter(buffer)
	defer tw.Close()

	return filepath.Walk(p, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fi, fi.Name())
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(strings.Replace(file, p, "", -1), string(filepath.Separator))
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			return err
		}

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		return nil
	})
}
