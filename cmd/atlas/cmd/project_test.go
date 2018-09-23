package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ahstn/atlas/pkg/builder/mocks"
	"github.com/ahstn/atlas/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
)

const (
	errCfgMissing = "config does not exist in ~/.config/atlas/ or pwd"
	errCfgFormat  = "config should be a .yaml file"
)

func TestPanicWithNoConfig(t *testing.T) {
	app := cli.App{
		Commands: []cli.Command{
			Project,
		},
	}

	panics := false
	var msg interface{}
	var err error
	func() {
		defer func() {
			if msg = recover(); msg != nil {
				panics = true
			}
		}()
		err = app.Run([]string{"foo", "project"})
	}()

	if err != nil {
		t.Fatal("Unexpected error occured:", err)
	}

	if !panics {
		t.Fatal("Expected a panic when no arguments are passed in")
	}

	if msg.(error).Error() != errCfgMissing {
		t.Fatal("Expected missing cfg message. Got:", msg)
	}
}

func TestPanicWithInvalidConfig(t *testing.T) {
	app := cli.App{
		Commands: []cli.Command{
			Project,
		},
	}

	cfg := []byte("---\nbase: ~/git\n")
	err := ioutil.WriteFile("./atlas.ya", cfg, 0644)
	if err != nil {
		t.Skip("'atlas.ya' already exists, skipping Test_PanicWithInvalidConfig", err)
	}

	panics := false
	var msg interface{}
	func() {
		defer func() {
			if msg = recover(); msg != nil {
				panics = true
			}
		}()
		err = app.Run([]string{"foo", "project", "-c", "atlas.ya"})
	}()

	if !panics {
		t.Fatal("Expected a panic when no arguments are passed in")
	}

	if msg.(error).Error() != errCfgFormat {
		t.Fatal("Expected missing cfg message. Got:", msg)
	}

	err = os.Remove("./atlas.ya")
	if err != nil {
		t.Skip("Can't Remove Dockerfile, skpping TestTryFindDockerfile")
	}
}

func TestRunAppBuild(t *testing.T) {
	app := config.Service{
		Tasks: []string{"clean", "build", "package"},
		Package: config.Package{
			SubDir: "package/",
		},
	}
	mvn := &mocks.Builder{}
	mvn.On("Run", mock.AnythingOfType("bool")).Return(nil)
	mvn.On("ModifyArgs", mock.AnythingOfType("[]string")).Return(nil)

	runAppBuild("", mvn, app, false)
	mvn.AssertNumberOfCalls(t, "ModifyArgs", 2)
	mvn.AssertNumberOfCalls(t, "Run", 2)

	app = config.Service{
		Tasks: []string{"clean", "build"},
	}
	mvn = &mocks.Builder{}
	mvn.On("Run", mock.AnythingOfType("bool")).Return(nil)

	runAppBuild("", mvn, app, false)
	mvn.AssertNumberOfCalls(t, "Run", 1)

	mvn = &mocks.Builder{}
	mvn.On("Run", mock.AnythingOfType("bool")).Return(errors.New("mock error"))

	err := runAppBuild("", mvn, app, false)
	assert.EqualError(t, err, "mock error")
}
