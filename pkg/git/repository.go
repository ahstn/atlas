package git

import (
	"os/exec"
	"path"
)

var execute = exec.Command

func gitCmdInDir(dir string, arg ...string) ([]byte, error) {
	arg = append([]string{"-C", dir}, arg...)
	cmd := execute("git", arg...)
	return cmd.CombinedOutput()
}

// Clone the passed repository into the path specified
func Clone(dir, repo, dest string) ([]byte, error) {
	cmd := execute("git", "clone", repo, path.Join(dir, dest))
	return cmd.CombinedOutput()
}

// CreateBranch creates and checks out a new branch
func CreateBranch(dir, branch string) ([]byte, error) {
	return gitCmdInDir(dir, "checkout", "-b", branch)
}

// CheckoutBranch switches branch
func CheckoutBranch(dir, branch string) ([]byte, error) {
	return gitCmdInDir(dir, "checkout", branch)
}

// Update the repository, pulling down all remote commits but keep local changes
func Update(dir string) ([]byte, error) {
	return gitCmdInDir(dir, "pull", "--rebase", "--prune")
}

// URL returns the remote repository URL
func URL(dir string) ([]byte, error) {
	return gitCmdInDir(dir, "ls-remote", "--get-url")
}

// Branch returns branch name the local repository is tracking to
func Branch(dir string) ([]byte, error) {
	return gitCmdInDir(dir, "rev-parse", "--abbrev-ref", "HEAD")
}
