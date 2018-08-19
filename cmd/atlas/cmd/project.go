package cmd

import (
	"context"

	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/ahstn/atlas/pkg/builder"
	"github.com/ahstn/atlas/pkg/config"
	"github.com/ahstn/atlas/pkg/docker"
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
		mvn.Dir = app.Path
		createAndRunBuilder(app.Path, mvn, *app, c)

		emoji.Printf(":whale:Building Dockerfile: %v [%v]...\n", app.Name, app.Path)
		if err != runDockerBuild(app.Path, *app) {
			panic(err)
		}
	}

	return nil
}

// TODO: Handle Package Args
func createAndRunBuilder(p string, mvn builder.Builder, app config.Service, c *cli.Context) {
	if app.HasTask("clean") {
		mvn.Clean()
	}
	if app.HasTask("build") {
		mvn.Build()
	}
	if !app.HasTask("test") {
		mvn.SkipTests()
	}
	if app.HasTask("package") && !app.HasPackageSubDir() {
		mvn.Package()
	}

	if err := mvn.Run(c.Bool("verbose")); err != nil {
		panic(err)
	}

	// In the event package pom lives in a seperate folder and needs to be ran
	// after the build, handle as such.
	if app.HasTask("package") && app.HasPackageSubDir() {
		mvn = &builder.Maven{Dir: p}

		mvn.Package()
		if err := mvn.Run(c.Bool("verbose")); err != nil {
			panic(err)
		}
	}
}

func runDockerBuild(p string, app config.Service) error {
	if !app.Docker.Enabled {
		emoji.Print("  :zzz: Docker build disabled. Skipping...\n")
		return nil
	}

	ctx := context.Background()
	return docker.ImageBuild(ctx, app.Docker)
}
