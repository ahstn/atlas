package git

import (
	"os/exec"
	"path"
)

//go:generate mockery -name SourceController -case underscore

// SourceController defines the actions that a VCS client should preform
type SourceController interface {
	Clone(string, string, string) ([]byte, error)
	CreateBranch(string, string) ([]byte, error)
	CheckoutBranch(string, string) ([]byte, error)
	Update(string) ([]byte, error)
	URL(string) ([]byte, error)
	Branch(string) ([]byte, error)
}

// Client is a implementation of SourceController for Git
type Client struct {
	cmd exec.Cmd
}

var execute = exec.Command

func gitCmdInDir(dir string, arg ...string) ([]byte, error) {
	arg = append([]string{"-C", dir}, arg...)
	cmd := execute("git", arg...)
	return cmd.CombinedOutput()
}

// Clone the passed repository into the path specified
func (c Client) Clone(dir, repo, dest string) ([]byte, error) {
	cmd := execute("git", "clone", repo, path.Join(dir, dest))
	return cmd.CombinedOutput()
}

// CreateBranch creates and checks out a new branch
func (c Client) CreateBranch(dir, branch string) ([]byte, error) {
	return gitCmdInDir(dir, "checkout", "-b", branch)
}

// CheckoutBranch switches branch
func (c Client) CheckoutBranch(dir, branch string) ([]byte, error) {
	return gitCmdInDir(dir, "checkout", branch)
}

// Update the repository, pulling down all remote commits but keep local changes
func (c Client) Update(dir string) ([]byte, error) {
	return gitCmdInDir(dir, "pull", "--rebase", "--prune")
}

// URL returns the remote repository URL
func (c Client) URL(dir string) ([]byte, error) {
	return gitCmdInDir(dir, "ls-remote", "--get-url")
}

// Branch returns branch name the local repository is tracking to
func (c Client) Branch(dir string) ([]byte, error) {
	return gitCmdInDir(dir, "rev-parse", "--abbrev-ref", "HEAD")
}
