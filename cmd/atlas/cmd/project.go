package cmd

import (
	"context"

	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/ahstn/atlas/pkg/builder"
	"github.com/ahstn/atlas/pkg/config"
	"github.com/ahstn/atlas/pkg/docker"
	"github.com/ahstn/atlas/pkg/util"
	"github.com/ahstn/atlas/pkg/validator"
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
	f, err := validator.ValidateExists(c.String("config"))
	if err != nil {
		panic(err)
	}

	if err = validator.ValidateConfig(f); err != nil {
		panic(err)
	}

	cfg, err := config.Read(f)
	if err != nil {
		panic(err)
	}

	mvn := &builder.Maven{}
	emoji.Printf(":file_folder:Operating in base directory [%v]\n", cfg.Root)
	for _, app := range cfg.Services {
		emoji.Printf("\n:wrench:Building: %v [%v]...\n", app.Name, app.Path)
		mvn = builder.NewClient(app.Path, app.Env, app.Tasks, app.Args)
		createAndRunBuilder(app.Path, mvn, *app, c)

		emoji.Printf(":whale:Building Dockerfile: %v [%v]...\n", app.Name, app.Path)
		if err != runDockerBuild(app.Path, *app) {
			panic(err)
		}
	}

	return runDockerLogs(ctx, cfg.Services)
}

// TODO: Handle Package Args
func createAndRunBuilder(p string, mvn builder.Builder, app config.Service, c *cli.Context) error {
	// In the event package pom lives in a seperate folder and needs to be ran
	// after the build, handle as such.
	if app.HasTask("package") && app.HasPackageSubDir() {
		mvn.ModifyArgs(util.StringSliceRemove(app.Tasks, "package"))
		if err := mvn.Run(c.Bool("verbose")); err != nil {
			return err
		}

		mvn.ModifyArgs([]string{"package", app.Package.Args})
		return mvn.Run(c.Bool("verbose"))
	}

	return mvn.Run(c.Bool("verbose"))
}

func runDockerBuild(p string, app config.Service) error {
	if !app.Docker.Enabled {
		emoji.Print("  :zzz: Docker build disabled. Skipping...\n")
		return nil
	}

	ctx := context.Background()
	return docker.ImageBuild(ctx, app.Docker)
}

func runDockerLogs(ctx context.Context, svcs []*config.Service) error {
	quit := make(chan bool)
	done := make(chan error)
	for _, app := range svcs {
		go func(d config.DockerArtifact) {
			var err error
			err = docker.RunContainer(ctx, d)

			select {
			case done <- err:
			case <-quit:
			}
		}(app.Docker)
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
