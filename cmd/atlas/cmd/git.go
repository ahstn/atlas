package cmd

import (
	"fmt"
	"path"

	"github.com/ahstn/atlas/cmd/atlas/flag"
	"github.com/ahstn/atlas/pkg/config"
	"github.com/ahstn/atlas/pkg/git"
	"github.com/ahstn/atlas/pkg/log"
	"github.com/ahstn/atlas/pkg/util"
	"github.com/ahstn/atlas/pkg/validator"
	"github.com/urfave/cli"
)

const (
	logOperating = ":file_folder:Operating in base directory [%s]\n"
	logSkipping  = "\n:zzz:Skipping Excluded App [%s]"
	logClone     = "\n:arrow_down:Cloning: %s [%s]..."
	logCheckout  = "\n:twisted_rightwards_arrows:Switching Branch: %s [%s]..."
	logNewBranch = "\n:recycle:Creating New Branch: %s [%s]..."
	logUpdating  = "\n:arrows_clockwise:Updating [%s]..."
)

var (
	// Git defines the command for the preforming git actions against services
	// i.e atlas git checkout master -e app1
	Git = cli.Command{
		Name:    "git",
		Aliases: []string{"g"},
		Usage:   "preform Git actions against service(s)",
		Subcommands: []cli.Command{
			branch,
			clone,
			checkout,
			update,
		},
	}

	clone = cli.Command{
		Name:      "clone",
		Usage:     "clone the services' repo(s) defined in config",
		ArgsUsage: "[single service]",
		Action:    gitCmd,
		Flags:     []cli.Flag{flag.Config, flag.Exclude},
	}

	branch = cli.Command{
		Name:      "branch",
		Usage:     "create a branch in the services' repo(s) defined in config",
		ArgsUsage: "[branch name]",
		Action:    gitCmd,
		Flags:     []cli.Flag{flag.Config, flag.Exclude},
	}

	checkout = cli.Command{
		Name:      "checkout",
		Usage:     "checkout a branch in services' repo(s) defined in config",
		ArgsUsage: "[branch]",
		Action:    gitCmd,
		Flags:     []cli.Flag{flag.Config, flag.Exclude},
	}

	update = cli.Command{
		Name:    "update",
		Aliases: []string{"up"},
		Usage:   "pull updates from remote, but keep local changes",
		Action:  gitCmd,
		Flags:   []cli.Flag{flag.Config, flag.Exclude},
	}
)

func gitWrapper(c *cli.Context) error {
	return gitCmd(c, log.NewClient(), new(git.Client))
}

func gitCmd(c *cli.Context, logger log.Logger, g git.SourceController) error {
	f, err := validator.ValidateExists(c.String("config"))
	logger.CheckError(err)

	cfg, err := config.Read(f)
	logger.CheckError(err)

	logger.Printf(logOperating, cfg.Root)
	for _, app := range cfg.Services {
		if util.StringSliceContains(c.StringSlice("exclude"), app.Name) {
			logger.Printf(logSkipping, app.Name)
			continue
		}

		var out []byte
		var err error
		switch c.Command.Name {
		case "clone":
			logger.Printf(logClone, app.Repo, app.Name)
			out, err = g.Clone(cfg.Root, app.Repo, app.Name)
		case "checkout":
			logger.Printf(logCheckout, c.Args().First(), app.Name)
			out, err = g.CheckoutBranch(path.Join(cfg.Root, app.Name), c.Args().First())
		case "branch":
			logger.Printf(logNewBranch, c.Args().First(), app.Name)
			out, err = g.CreateBranch(path.Join(cfg.Root, app.Name), c.Args().First())
		case "update":
			logger.Printf(logUpdating, app.Name)
			out, err = g.Update(path.Join(cfg.Root, app.Name))
		}

		logger.CheckError(err)
		fmt.Print("\n\t", string(out))
	}

	return nil
}
