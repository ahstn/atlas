package cmd

import (
	"flag"
	"testing"

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

	url, err := determineURL(c)
	if err != nil {
		t.Fatal("Failed to determine URL:", err)
	}

	if url != expectedURL {
		t.Fatalf("Expected determineURL to return '%s'. Got: %v", expectedURL, err)
	}
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
