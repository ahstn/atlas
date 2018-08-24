package main

import (
	"log"
	"os"

	"github.com/ahstn/atlas/cmd/atlas/cmd"
	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:     "atlas",
		Usage:    "Make Development Great Again",
		Version:  "0.1.0-beta-3",
		HelpName: "atlas",
		Commands: []cli.Command{
			cmd.Build,
			cmd.Docker,
			cmd.Git,
			cmd.Project,
			cmd.Repo,
			cmd.Issues,
		},
		Flags: []cli.Flag{
			flag.SkipTests,
			flag.Verbose,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
