package cmd

import (
	"testing"

	"github.com/urfave/cli"
)

const (
	errArg      = "invalid path argument - you must pass a directory"
	errTag      = "invalid tag - you must specify a version"
	errBuildArg = "arg has no value - must be in `key=value` format"
)

func Test_PanicWithNoArguments(t *testing.T) {
	app := cli.App{
		Commands: []cli.Command{
			Docker,
		},
	}

	panics := false
	var msg interface{}
	func() {
		defer func() {
			if msg = recover(); msg != nil {
				panics = true
			}
		}()
		app.Run([]string{"foo", "docker"})
	}()

	if !panics {
		t.Fatal("Expected a panic when no arguments are passed in")
	}

	if msg.(error).Error() != errArg {
		t.Fatal("Expected invalid args message. Got:", msg)
	}
}

func Test_PanicWithInvalidTag(t *testing.T) {
	app := cli.App{
		Commands: []cli.Command{
			Docker,
		},
	}

	panics := false
	var msg interface{}
	func() {
		defer func() {
			if msg = recover(); msg != nil {
				panics = true
			}
		}()
		app.Run([]string{"foo", "docker", "./", "-t", "ahstn"})
	}()

	if !panics {
		t.Fatal("Expected a panic when an invalid tag is passed in")
	}

	if msg.(error).Error() != errTag {
		t.Fatal("Expected invalid args message. Got:", msg)
	}
}

func Test_PanicWithInvalidBuildArg(t *testing.T) {
	app := cli.App{
		Commands: []cli.Command{
			Docker,
		},
	}

	panics := false
	var msg interface{}
	func() {
		defer func() {
			if msg = recover(); msg != nil {
				panics = true
			}
		}()
		app.Run([]string{"foo", "docker", "./", "-t", "ahstn:1.0", "-a", "VERSION"})
	}()

	if !panics {
		t.Fatal("Expected a panic when invalid build args are passed in")
	}

	if msg.(error).Error() != errBuildArg {
		t.Fatal("Expected invalid args message. Got:", msg)
	}
}
