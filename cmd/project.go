package cmd

import (
	"log"
	"os"
	"path"

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
		flag.Verbose,
	},
}

// ProjectAction executes the logic to read a project file and build it's apps
func ProjectAction(c *cli.Context) error {
	h := path.Join(os.Getenv("HOME"), ".config/atlas/")
	p, err := config.Read(path.Join(h, c.String("config")))
	if err != nil {
		log.Printf("File not found. Error: %v", err)
		// TODO: try './atlas.yaml'
	}

	for _, app := range p.Services {
		log.Printf("Service: %v", app.Name)
	}

	log.Printf("Operating in base directory: %v", p.Root)

	return nil
}
