package main

import (
	"log"
	"os"

	"github.com/ahstn/atlas/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		cmd.Clean,
		cmd.Build,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
