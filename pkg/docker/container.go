package docker

import (
	"strings"

	"github.com/ahstn/atlas/pkg/config"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
)

// RunContainer takes a DockerArtifact and runs it using the Docker Daemon
func RunContainer(c context.Context, d config.DockerArtifact) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

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

	if err := cli.ContainerStart(c, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	out, err := cli.ContainerLogs(c, resp.ID, types.ContainerLogsOptions{ShowStdout: true, Follow: true})
	if err != nil {
		return err
	}

	PrintRun(out, strings.Split(d.Tag, ":")[0])

	statusCh, errCh := cli.ContainerWait(c, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	return nil
}

func NewClient() (*client.Client, error) {
	return client.NewEnvClient()
}
