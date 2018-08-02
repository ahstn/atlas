package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	gitCheckoutResult     = "Switched to branch 'master'"
	gitCreateBranchResult = "Switched to new branch 'feature/testing'"
)

func TestCheckout(t *testing.T) {
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := CheckoutBranch("/tmp/", "master")
	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}
	if string(out) != gitCheckoutResult {
		t.Errorf("Expected %q, got %q", gitCheckoutResult, out)
	}
}

func TestCreateBranch(t *testing.T) {
	execute = fakeExecCommand
	defer func() { execute = exec.Command }()
	out, err := CreateBranch("/tmp/", "feature/testing")
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
	case strings.Contains(strings.Join(args, " "), "checkout -b"):
		fmt.Fprintf(os.Stdout, gitCreateBranchResult)
	case strings.Contains(strings.Join(args, " "), "checkout"):
		fmt.Fprintf(os.Stdout, gitCheckoutResult)
	}

	os.Exit(0)
}
