package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ahstn/atlas/builder"
	"github.com/ahstn/atlas/config"
	"github.com/ahstn/atlas/flag"
	"github.com/urfave/cli"
)

// Project defines the command for the cli to process an entire project
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
func ProjectAction(c *cli.Context) error {
	h := path.Join(os.Getenv("HOME"), ".config/atlas/")
	p, err := config.Read(path.Join(h, c.String("config")))
	if err != nil {
		log.Printf("File not found. Error: %v", err)
	}

	fmt.Println("Operating in base directory: " + p.Root)
	for _, app := range p.Services {
		fmt.Println("\nBuilding: " + app.Name)
		createAndRunBuilder(path.Join(p.Root, app.Name), app, c)
	}

	return nil
}

func createAndRunBuilder(p string, app config.Service, c *cli.Context) {
	mvn := builder.Maven{
		Dir: p,
	}

	if c.Bool("clean") {
		mvn.Clean()
	}
	if c.Bool("skipTests") {
		mvn.SkipTests()
	}

	// In the event package pom lives in a seperate folder and needs to be ran
	// after the build, handle as such.
	if app.HasTask("package") && app.HasPackageSubDir() {
		packageMvn := mvn
		mvn.Build()
		if err := mvn.Run(c.Bool("verbose")); err != nil {
			log.Printf("Error:" + err.Error())
			os.Exit(1)
		}

		packageMvn.Package()
		if err := packageMvn.Run(c.Bool("verbose")); err != nil {
			log.Printf("Error:" + err.Error())
			os.Exit(1)
		}
	} else {
		mvn.Build()
		if err := mvn.Run(c.Bool("verbose")); err != nil {
			log.Printf("Error:" + err.Error())
			os.Exit(1)
		}
	}
}
