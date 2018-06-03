package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
	"gopkg.in/cheggaaa/pb.v2"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

func main() {
	about := emoji.Sprint("A collection of development workflow tools :wrench:")

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

	bar := fmt.Sprintf(
		`{{"%s"}} {{bar . "|" "██" "░" "░" "|" | green}} {{speed . | blue }}`,
		emoji.Sprint(" :wrench:Building"),
	)

	runProgressBar(bar)
}

func runProgressBar(tmpl string) {
	count := 1000
	bar := pb.ProgressBarTemplate(tmpl).Start(count)
	bar.Set("prefix", "Testing..")
	bar.SetWidth(80)
	defer bar.Finish()
	for i := 0; i < count/2; i++ {
		bar.Add(2)
		time.Sleep(time.Millisecond * 4)
	}
}
