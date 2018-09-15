package docker

import (
	"fmt"
	"strings"

	"github.com/ahstn/atlas/pkg/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

// RunContainer takes a DockerArtifact and runs it using the Docker Daemon
func RunContainer(c context.Context, cli *client.Client, d *config.DockerArtifact) error {
	ports, portBindings, err := nat.ParsePortSpecs(d.Ports)
	if err != nil {
		return err
	}

	resp, err := cli.ContainerCreate(
		c,
		&container.Config{
			Image:        d.Tag,
			Cmd:          strings.Split(d.Cmd, " "),
			Env:          d.Env,
			ExposedPorts: ports,
			Tty:          true,
		},
		&container.HostConfig{
			PortBindings: portBindings,
		},
		nil,
		"",
	)
	if err != nil {
		return err
	}
	d.ID = resp.ID

	if err := cli.ContainerStart(c, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	out, err := cli.ContainerLogs(c, resp.ID, types.ContainerLogsOptions{ShowStdout: true, Follow: true})
	if err != nil {
		return err
	}

	PrintRun(out, strings.Split(d.Tag, ":")[0])

	return nil
}

// StopAndRemoveContainer stops and removes the container `c` using the `cli`
func StopAndRemoveContainer(ctx context.Context, cli *client.Client, d config.DockerArtifact) error {
	if d.ID == "" {
		fmt.Println("NO ID")
		return nil // if the ID isn't set, the container probably wasn't started
	}

	err := cli.ContainerStop(ctx, d.ID, nil)
	if err != nil {
		return err
	}

	err = cli.ContainerRemove(ctx, d.ID, types.ContainerRemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}
