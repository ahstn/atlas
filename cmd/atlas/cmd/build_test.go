package cmd

import (
	"flag"
	"testing"

	"github.com/ahstn/atlas/pkg/builder/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
)

func TestCleanBuild(t *testing.T) {
	c := setupContext()
	c.Set("clean", "true")

	mvn := &mocks.BuilderPartial{}
	mvn.On("Run", mock.AnythingOfType("bool")).Return(nil)

	build(c, mvn)
	assert.Equal(t, "mvn --batch-mode clean install", mvn.Args())
}

func TestBuild(t *testing.T) {
	c := setupContext()
	mvn := &mocks.BuilderPartial{}
	mvn.On("Run", mock.AnythingOfType("bool")).Return(nil)

	build(c, mvn)
	assert.Equal(t, "mvn --batch-mode install", mvn.Args())
}

func setupContext() *cli.Context {
	flags := flag.NewFlagSet("test", 0)
	flags.Bool("clean", false, "usage")
	return cli.NewContext(nil, flags, nil)
}
