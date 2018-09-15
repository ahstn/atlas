package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	gitCloneResult        = "Cloning into 'atlas'..."
	gitCheckoutResult     = "Switched to branch 'master'"
	gitCreateBranchResult = "Switched to new branch 'feature/testing'"
	gitUpdateResult       = "Current branch master is up to date."
)

func TestClone(t *testing.T) {
	g := new(Client)
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := g.Clone("/tmp/", "https://github.com/ahstn/atlas", "atlas-http")
	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}
	if string(out) != gitCloneResult {
		t.Errorf("Expected %q, got %q", gitCloneResult, out)
	}
}

func TestCheckout(t *testing.T) {
	g := new(Client)
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := g.CheckoutBranch("/tmp/", "master")
	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}
	if string(out) != gitCheckoutResult {
		t.Errorf("Expected %q, got %q", gitCheckoutResult, out)
	}
}

func TestCreateBranch(t *testing.T) {
	g := new(Client)
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := g.CreateBranch("/tmp/", "feature/testing")
	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}
	if string(out) != gitCreateBranchResult {
		t.Errorf("Expected %q, got %q", gitCreateBranchResult, out)
	}
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
	case strings.Contains(strings.Join(args, " "), "checkout"):
		fmt.Fprintf(os.Stdout, gitCheckoutResult)
	}

	os.Exit(0)
}
