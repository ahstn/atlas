package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/ahstn/atlas/pkg/builder"
	"github.com/ahstn/atlas/pkg/config"
	"github.com/ahstn/atlas/pkg/docker"
	"github.com/ahstn/atlas/pkg/util"
	"github.com/docker/docker/client"
	"github.com/urfave/cli"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

// Project defines the command for the cli to process an entire project
// utilising an atlas.yaml config file
var Project = cli.Command{
	Name:    "project",
	Aliases: []string{"p"},
	Usage:   "build project (collection of services)",
	Action:  ProjectAction,
	Flags: []cli.Flag{
		flag.Config,
		flag.SkipTests,
		flag.Verbose,
	},
}

// ProjectAction executes the logic to read a project file and build it's apps
// TODO: Proper logging and error handling
func ProjectAction(c *cli.Context) error {
	ctx := context.Background()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	cfg, err := config.ReadAndValidate(c.String("config"))
	if err != nil {
		panic(err)
	}

	mvn := &builder.Maven{}
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	go DetectShutdown(ctx, cli, cfg, quit)

	emoji.Printf(":file_folder:Operating in base directory [%v]\n", cfg.Root)
	for _, app := range cfg.Services {
		emoji.Printf("\n:wrench:Building: %v [%v]...\n", app.Name, app.Path)
		mvn = builder.NewClient(app.Path, app.Env, app.Tasks, app.Args)
		if err != runAppBuild(app.Path, mvn, *app, c.Bool("verbose")) {
			panic(err)
		}

		emoji.Printf(":whale:Building Dockerfile: %v [%v]...\n", app.Name, app.Path)
		if err != runDockerBuild(cli, app.Path, *app) {
			panic(err)
		}
	}

	return runDockerLogs(ctx, cfg.Services)
}

// TODO: Handle Package Args
func runAppBuild(p string, mvn builder.Builder, app config.Service, v bool) error {
	if len(app.Tasks) == 0 {
		emoji.Print("  :zzz: Maven build disabled. Skipping...\n")
		return nil
	}

	// In the event package pom lives in a seperate folder and needs to be ran
	// after the build, handle as such.
	if app.HasTask("package") && app.HasPackageSubDir() {
		mvn.ModifyArgs(util.StringSliceRemove(app.Tasks, "package"))
		if err := mvn.Run(v); err != nil {
			return err
		}

		mvn.ModifyArgs([]string{"package", app.Package.Args})
		return mvn.Run(v)
	}

	return mvn.Run(v)
}

func runDockerBuild(cli *client.Client, p string, app config.Service) error {
	if !app.Docker.Enabled {
		emoji.Print("  :zzz: Docker build disabled. Skipping...\n")
		return nil
	}

	ctx := context.Background()
	return docker.ImageBuild(ctx, cli, app.Docker)
}

func runDockerLogs(ctx context.Context, svcs []*config.Service) error {
	quit := make(chan bool)
	done := make(chan error)
	for _, app := range svcs {
		go func(d *config.DockerArtifact) {
			var err error
			cli, err := client.NewEnvClient()
			err = docker.RunContainer(ctx, cli, d)

			select {
			case done <- err:
			case <-quit:
			}
		}(&app.Docker)
	}

	for range svcs {
		err := <-done
		if err != nil {
			close(quit)
			return err
		}
	}

	return nil
}

// DetectShutdown actives on SIGTERM when the user cancels the process and
// gracefully shuts down all running containers
func DetectShutdown(ctx context.Context, cli *client.Client, cfg config.Project, quit chan os.Signal) {
	<-quit
	fmt.Println("\n\nUser Shutdown - Cleaning up containers..")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for _, app := range cfg.Services {
		docker.StopAndRemoveContainer(ctx, cli, app.Docker)
	}
	os.Exit(1)
}
