package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

func main() {
	about := emoji.Sprint(":wrench: A collection of development workflow tools")

	app := cli.NewApp()
	app.Name = "atlas"
	app.Usage = about
	app.Action = func(c *cli.Context) error {
		fmt.Println(about)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
