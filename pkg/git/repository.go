package git

import (
	"os/exec"
	"path"
)

var execute = exec.Command

func gitCmdInDir(dir string, arg ...string) error {
	arg = append([]string{"-C", dir}, arg...)
	cmd := execute("git", arg...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

// Clone the passed repository into the path specified
func Clone(dir, repo, dest string) error {
	cmd := execute("git", "clone", repo, path.Join(dir, dest))
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

// CreateBranch creates and checks out a new branch
func CreateBranch(dir, branch string) error {
	return gitCmdInDir(dir, "checkout", "-b", branch)
}

// CheckoutBranch switches branch
func CheckoutBranch(dir, branch string) error {
	return gitCmdInDir(dir, "checkout", branch)
}

// Update the repository, pulling down all remote commits but keep local changes
func Update(dir string) error {
	return gitCmdInDir(dir, "pull", "--rebase", "--prune")
}
