package main

import (
	"log"
	"os"

	"github.com/ahstn/atlas/cmd"
	"github.com/ahstn/atlas/flag"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "atlas",
		Usage: "Make Development Great Again",
		Commands: []cli.Command{
			cmd.Build,
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
