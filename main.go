package main

import (
	"log"
	"os"

	"github.com/ahstn/atlas/cmd"
	"github.com/ahstn/atlas/flag"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		cmd.Clean,
		cmd.Build,
		cmd.Package,
	}
	app.Flags = []cli.Flag{
		flag.SkipTests,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
