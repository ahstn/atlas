package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/urfave/cli"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

// Repo defines the command for the cli to open a repo's URL
var Repo = cli.Command{
	Name:    "repo",
	Aliases: []string{"r"},
	Usage:   "Open Git repo in browser",
	Action:  RepoAction,
}

// RepoAction executes the logic to open a repo's URL
func RepoAction(c *cli.Context) error {
	cmd := exec.Command("git", "ls-remote", "--get-url")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	url, err := processRepoURL(string(out))
	if err != nil {
		panic(err)
	}

	emoji.Printf(":globe_with_meridians:Opening Repo URL: %v \n", url)
	openBrowser(url)
	if err != nil {
		panic(err)
	}
	return nil
}

func processRepoURL(r string) (string, error) {
	if strings.Contains(r, "git@") {
		r = strings.Replace(r, "git@", "https://", 1)
		r = strings.Replace(r, ".com:", ".com/", 1)

		return r, nil
	} else if strings.Contains(r, "https://") {
		return r, nil
	}

	return "", errors.New("could not process git repo url")
}

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
