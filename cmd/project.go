package cmd

import (
	"fmt"
	"path"

	"github.com/ahstn/atlas/builder"
	"github.com/ahstn/atlas/config"
	"github.com/ahstn/atlas/flag"
	"github.com/urfave/cli"
)

// Project defines the command for the cli to process an entire project
// utilising an atlas.yaml config file
var Project = cli.Command{
	Name:    "project",
	Aliases: []string{"p"},
	Usage:   "Build Project (Collection of Services)",
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
	f, err := config.ValidateExists(c.String("config"))
	if err != nil {
		panic(err)
	}

	if err = config.ValidateConfig(f); err != nil {
		panic(err)
	}

	cfg, err := config.Read(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Operating in base directory: [%v]\n", cfg.Root)
	for _, app := range cfg.Services {
		fmt.Printf("\nBuilding: %v [%v]\n", app.Name, path.Join(cfg.Root, app.Name))
		createAndRunBuilder(path.Join(cfg.Root, app.Name), app, c)
	}

	return nil
}

// TODO: Handle Package Args
func createAndRunBuilder(p string, app config.Service, c *cli.Context) {
	mvn := builder.Maven{Dir: p}

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
		mvn = builder.Maven{Dir: p}

		mvn.Package()
		if err := mvn.Run(c.Bool("verbose")); err != nil {
			panic(err)
		}
	}
}
