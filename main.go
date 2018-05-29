package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"

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

	p := mpb.New(
		mpb.WithWidth(48),
		mpb.WithFormat("|██░|"),
		//mpb.WithFormat("|█ |"),
		mpb.WithRefreshRate(180*time.Millisecond),
	)

	total := 100
	build := emoji.Sprint(":wrench: Building")
	// adding a single bar
	bar := p.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.StaticName(build, len(build)+1, decor.DSyncSpace),
		),
		mpb.PrependDecorators(
			decor.CountersNoUnit("%d / %d", 12, 0),
		),
		mpb.AppendDecorators(
			decor.Percentage(5, 0),
		),
	)

	// simulating some work
	max := 100 * time.Millisecond
	for i := 0; i < total; i++ {
		time.Sleep(time.Duration(rand.Intn(10)+1) * max / 10)
		bar.Increment()
	}
	p.Wait()
}
