package cmd

import (
	"flag"
	"os"
	"testing"

	"github.com/ahstn/atlas/pkg/git"
	"github.com/ahstn/atlas/pkg/git/mocks"
	logMocks "github.com/ahstn/atlas/pkg/log/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
)

const (
	testRepo    = "https://github.com/ahstn/atlas"
	expectedURL = "https://github.com/ahstn/atlas/issues"
)

func TestDetermineURLUsingFlag(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	set.String("url", "", "usage")
	c := cli.NewContext(nil, set, nil)

	// Emulate CLI command
	c.Set("url", testRepo)

	url, err := determineURL(c, new(git.Client))
	assert.NoError(t, err)

	if url != expectedURL {
		t.Fatalf("Expected determineURL to return '%s'. Got: %v", expectedURL, err)
	}
}

func TestIssues(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	set.String("url", "", "usage")
	c := cli.NewContext(nil, set, nil)

	// Emulate CLI command
	c.Set("url", testRepo)

	logger := &logMocks.Logger{}
	logger.On("CheckError", nil).Return(nil)
	logger.On("Printf", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	git := &mocks.SourceController{}
	git.On("Branch", mock.AnythingOfType("string")).Return([]byte("master"), nil)
	git.On("URL", mock.AnythingOfType("string")).Return(testRepo)

	os.Setenv("ATLAS_UNIT_TEST", "1") // Prevent link actually being opened
	issues(c, logger, git)
	logger.AssertNumberOfCalls(t, "CheckError", 2)
}

// Even though this relies on Git, I'm happy keeping it as the remote URL should
// always be the GitHub URL for ahstn/atlas
// NB: Disabled for now as CircleCI build agent doesn't have Git
// func TestDetermineURLUsingGit(t *testing.T) {
// 	set := flag.NewFlagSet("test", 0)
// 	c := cli.NewContext(nil, set, nil)

// 	url, err := determineURL(c)
// 	if err != nil {
// 		t.Fatal("Failed to determine URL:", err)
// 	}

// 	if url != expectedURL {
// 		t.Fatalf("Expected determineURL to return '%s'. Got: %v", expectedURL, err)
// 	}
// }
