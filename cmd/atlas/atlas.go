package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahstn/atlas/cmd/atlas/cmd"
	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:     "atlas",
		Usage:    "Make Development Great Again",
		Version:  "0.1.0",
		HelpName: "atlas",
		Commands: []cli.Command{
			cmd.Build,
			cmd.Docker,
			cmd.Git,
			cmd.Issues,
			cmd.Project,
			cmd.Repo,
		},
		Flags: []cli.Flag{
			flag.SkipTests,
			flag.Verbose,
		},
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nUser Shutdown..")
		os.Exit(1)
	}()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
