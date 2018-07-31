package git

import (
	"fmt"
	"os"
	"os/exec"
)

var gitCmd = exec.Command

// Clone the passed repository into the path specified
func Clone(dir, repo, dest string) error {
	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	cmd := gitCmd("git", "clone", repo, dest)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(out)
	return nil
}

// CreateBranch creates and checks out a new branch
func CreateBranch(dir, branch string) error {
	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	cmd := gitCmd("git", "checkout", "-b", branch)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(out)
	return nil
}

// CheckoutBranch switches branch
func CheckoutBranch(dir, branch string) error {
	fmt.Println(dir)
	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	cmd := gitCmd("git", "checkout", branch)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(out)
	return nil

}

// Update the repository, pulling down all remote commits but keep local changes
func Update(dir string) error {
	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	cmd := gitCmd("git", "pull", "--rebase", "--prune")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(out)
	return nil
}
