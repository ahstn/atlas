package cmd

import (
	"path"

	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/ahstn/atlas/pkg/config"
	"github.com/ahstn/atlas/pkg/git"
	"github.com/ahstn/atlas/pkg/validator"
	"github.com/urfave/cli"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

// Git defines the command for the preforming git actions against services
// i.e atlas git checkout master -e app1
var Git = cli.Command{
	Name:    "git",
	Aliases: []string{"g"},
	Usage:   "preform Git actions against service(s)",
	Subcommands: []cli.Command{
		{
			Name:      "clone",
			Usage:     "clone the services' repo(s) defined in config",
			ArgsUsage: "[single service]",
			Action:    cloneAction,
			Flags:     []cli.Flag{flag.Config},
		},
		{
			Name:      "branch",
			Usage:     "create a branch in the services' repo(s) defined in config",
			ArgsUsage: "[branch name]",
			Action:    branchAction,
			Flags:     []cli.Flag{flag.Config},
		},
		{
			Name:      "checkout",
			Usage:     "checkout a branch in services' repo(s) defined in config",
			ArgsUsage: "[branch]",
			Action:    checkoutAction,
			Flags:     []cli.Flag{flag.Config},
		},
		{
			Name:    "update",
			Aliases: []string{"up"},
			Usage:   "pull updates from remote, but keep local changes",
			Flags:   []cli.Flag{flag.Config},
		},
	},
	Flags: []cli.Flag{
    flag.Exclude
		flag.Config,
	},
}

func cloneAction(c *cli.Context) error {
	f, err := validator.ValidateExists(c.String("config"))
	if err != nil {
		panic(err)
	}

	cfg, err := config.Read(f)
	if err != nil {
		panic(err)
	}

	emoji.Printf(":file_folder:Operating in base directory [%v]\n", cfg.Root)
	for _, app := range cfg.Services {
		emoji.Printf("\n:arrow_down:Cloning: %v [%v]...\n", app.Name, app.Repo)

		_, err := git.Clone(cfg.Root, app.Repo, app.Name)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func checkoutAction(c *cli.Context) error {
	f, err := validator.ValidateExists(c.String("config"))
	if err != nil {
		panic(err)
	}

	cfg, err := config.Read(f)
	if err != nil {
		panic(err)
	}

	emoji.Printf(":file_folder:Operating in base directory [%v]", cfg.Root)
	for _, app := range cfg.Services {
		emoji.Printf("\n:arrow_down:Cloning: %v [%v]...\n", app.Name, app.Repo)
		git.CheckoutBranch(path.Join(cfg.Root, app.Name), c.Args().First())
	}

	return nil
}

func branchAction(c *cli.Context) error {
	f, err := validator.ValidateExists(c.String("config"))
	if err != nil {
		panic(err)
	}

	cfg, err := config.Read(f)
	if err != nil {
		panic(err)
	}

	emoji.Printf(":file_folder:Operating in base directory [%v]", cfg.Root)
	for _, app := range cfg.Services {
		emoji.Printf("\n:arrow_down:Cloning: %v [%v]...\n", app.Name, app.Repo)
		git.CreateBranch(path.Join(cfg.Root, app.Name), c.Args().First())
	}

	return nil
}

func updateAction(c *cli.Context) error {
	f, err := validator.ValidateExists(c.String("config"))
	if err != nil {
		panic(err)
	}

	cfg, err := config.Read(f)
	if err != nil {
		panic(err)
	}

	emoji.Printf(":file_folder:Operating in base directory [%v]", cfg.Root)
	for _, app := range cfg.Services {
		emoji.Printf("\n:arrow_down:Updating: %v [%v]...\n", app.Name, app.Repo)
		git.Update(path.Join(cfg.Root, app.Name))
	}

	return nil
}
