package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/urfave/cli"
)

const (
	errCfgMissing = "config does not exist in ~/.config/atlas/ or pwd"
	errCfgFormat  = "config should be a .yaml file"
)

func Test_PanicWithNoConfig(t *testing.T) {
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

func Test_PanicWithInvalidConfig(t *testing.T) {
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

// func Test_createAndRunBuilder(t *testing.T) {
// 	mvn := &mocks.Builder{}
// 	app := config.Service{
// 		Tasks: []string{"clean", "build"},
// 	}

// 	set := flag.NewFlagSet("test", 0)
// 	set.Bool("verbose", false, "doc")
// 	globalSet := flag.NewFlagSet("test", 0)
// 	globalCtx := cli.NewContext(nil, globalSet, nil)
// 	c := cli.NewContext(nil, set, globalCtx)

// 	mvn.On("Clean").Return(nil)
// 	mvn.On("Build").Return(nil)
// 	mvn.On("SkipTests").Return(nil)
// 	mvn.On("Run", mock.AnythingOfType("bool")).Return(nil)
// 	createAndRunBuilder("", mvn, app, c)

// 	if !strings.Contains(mvn.Args(), "clean") {
// 		t.Fatal("Expected args to include 'clean'. Got:", mvn.Args())
// 	}

// 	if !strings.Contains(mvn.Args(), "build") {
// 		t.Fatal("Expected args to include 'build'. Got:", mvn.Args())
// 	}

// 	if !strings.Contains(mvn.Args(), "-DskipTests") {
// 		t.Fatal("Expected args to include '-DskipTests'. Got:", mvn.Args())
// 	}
// }
