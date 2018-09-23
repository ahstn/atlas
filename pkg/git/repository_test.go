package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	gitCloneResult        = "Cloning into 'atlas'..."
	gitCheckoutResult     = "Switched to branch 'master'"
	gitCreateBranchResult = "Switched to new branch 'feature/testing'"
	gitUpdateResult       = "Current branch master is up to date."
	gitURLResult          = "git@github.com:ahstn/atlas.git"
	gitBranchResult       = "master"
)

func TestClone(t *testing.T) {
	g := new(Client)
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := g.Clone("/tmp/", "https://github.com/ahstn/atlas", "atlas-http")
	assert.NoError(t, err)
	assert.Equal(t, gitCloneResult, string(out))
}

func TestCheckout(t *testing.T) {
	g := new(Client)
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := g.CheckoutBranch("/tmp/", "master")

	assert.NoError(t, err)
	assert.Equal(t, gitCheckoutResult, string(out))
}

func TestCreateBranch(t *testing.T) {
	g := new(Client)
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := g.CreateBranch("/tmp/", "feature/testing")

	assert.NoError(t, err)
	assert.Equal(t, gitCreateBranchResult, string(out))
}

func TestUpdateBranch(t *testing.T) {
	g := new(Client)
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := g.Update("/tmp/")

	assert.NoError(t, err)
	assert.Equal(t, gitUpdateResult, string(out))
}

func TestURL(t *testing.T) {
	g := new(Client)
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := g.URL("/tmp/")

	assert.NoError(t, err)
	assert.Equal(t, gitURLResult, string(out))
}

func TestBranch(t *testing.T) {
	g := new(Client)
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := g.Branch("/tmp/")

	assert.NoError(t, err)
	assert.Equal(t, gitBranchResult, string(out))
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

// TestHelperProcess isn't a real test. It's used as a helper process
// for fakeExecCommand.
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	cmd, args := os.Args[0], os.Args[1:]
	if !strings.Contains(cmd, "git") {
		fmt.Fprintf(os.Stderr, "Expected command to be 'git'. Got: '%s' %s", cmd, args)
		os.Exit(2)
	}

	// TODO: Futher Validation
	switch {
	case strings.Contains(strings.Join(args, " "), "clone"):
		fmt.Fprintf(os.Stdout, gitCloneResult)
	case strings.Contains(strings.Join(args, " "), "pull"):
		if !strings.Contains(strings.Join(args, " "), "--rebase") {
			fmt.Fprintf(os.Stderr, "Expected '--rebase' flag. Got: %s", args)
		} else if !strings.Contains(strings.Join(args, " "), "--prune") {
			fmt.Fprintf(os.Stderr, "Expected '--prune' flag. Got: %s", args)
		}
		fmt.Fprintf(os.Stdout, gitUpdateResult)
	case strings.Contains(strings.Join(args, " "), "checkout -b"):
		fmt.Fprintf(os.Stdout, gitCreateBranchResult)
	case argsMatchCommand(args, "checkout"):
		fmt.Fprintf(os.Stdout, gitCheckoutResult)
	case argsMatchCommand(args, "ls-remote"):
		fmt.Fprintf(os.Stdout, gitURLResult)
	case argsMatchCommand(args, "rev-parse"):
		fmt.Fprintf(os.Stdout, gitBranchResult)
	}

	os.Exit(0)
}

func argsMatchCommand(args []string, cmd string) bool {
	return strings.Contains(strings.Join(args, " "), cmd)
}
