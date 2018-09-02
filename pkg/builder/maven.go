package builder

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/ahstn/atlas/pkg/pb"
	"github.com/ahstn/atlas/pkg/util"
	"github.com/briandowns/spinner"
)

// Maven is a implmentation of Builder{} for Java
type Maven struct {
	Dir string
	cmd exec.Cmd
	out io.Writer
}

// NewClient initialises a new Maven client based on the parameters passed
// Goals are the tasks Maven should preform (i.e. install, package, etc).
// Args are variables passed to the Maven build (i.e. -DskipTests).
// --batch-mode is used to remove extra escape characters for color and bold.
func NewClient(dir string, env, goals, args []string) *Maven {
	args = append(goals, args...)
	path, _ := exec.LookPath("mvn")
	return &Maven{
		Dir: dir,
		cmd: exec.Cmd{
			Path: path,
			Args: append([]string{"mvn", "--batch-mode"}, args...),
			Env:  env,
			Dir:  dir,
		},
		out: os.Stdout,
	}
}

// NewCustomClient initialises a new Maven client based on the parameters passed.
// With the added option to override the Maven binary path
func NewCustomClient(path, dir string, env, goals, args []string) Maven {
	args = append(goals, args...)
	return Maven{
		cmd: exec.Cmd{
			Path: path,
			Args: append([]string{"mvn", "--batch-mode"}, args...),
			Env:  env,
		},
		out: os.Stdout,
	}
}

// Run executes the built command
func (m Maven) Run(v bool) error {
	stdoutPipe, err := m.cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
	}

	scanner := bufio.NewScanner(stdoutPipe)
	var wg sync.WaitGroup
	if v {
		go printVerboseLog(scanner, m.out)
	} else {
		wg.Add(1)
		go printLog(scanner, m.out, &wg)
	}

	err = m.cmd.Start()
	if err != nil {
		return err
	}

	err = m.cmd.Wait()
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func checkError(s string, out io.Writer, queue []*spinner.Spinner) {
	var ignored = []string{
		"To see the full", "Re-run Maven using", "For more info",
	}
	if strings.Contains(s, "ERROR") {
		if strings.TrimSpace(s) != "[ERROR]" && !util.StringContainsAny(s, ignored...) {
			x := queue[0]
			x.Stop()
			fmt.Fprintln(out, s)
		}
	}
}

func printLog(s *bufio.Scanner, out io.Writer, wg *sync.WaitGroup) {
	failedTests := false
	queue := make([]*spinner.Spinner, 0)
	//TODO: Handle Packaging output
	for s.Scan() {
		checkError(s.Text(), out, queue)
		if strings.Contains(s.Text(), "Failed tests:") {
			failedTests = true
		} else if strings.Contains(s.Text(), "Tests run:") {
			failedTests = false
		}

		if failedTests {
			fmt.Fprintf(out, "\n%s", s.Text())
		} else if len(s.Text()) > 30 && strings.Contains(s.Text()[:30], "Building") {
			// If another "Building" string is detected, last build is done
			// therefore update the last spinner
			if len(queue) != 0 {
				x := queue[0]
				x.Stop()
				queue = queue[1:]
			}

			module := strings.SplitAfter(s.Text(), "Building")[1]

			// Replacing full path with project's build directory
			if strings.Contains(module, ".jar") {
				dir := path.Base(os.Getenv("PWD"))
				module = strings.Replace(module, os.Getenv("PWD"), dir, 1)
			}

			// Create spinner and add it to the queue of pending builds
			spinner := pb.Fprint(out, module)
			queue = append(queue, spinner)
		}
	}

	for _, x := range queue {
		x.Stop()
		queue = queue[1:]
	}

	wg.Done()
}

func printVerboseLog(scanner *bufio.Scanner, out io.Writer) {
	for scanner.Scan() {
		fmt.Fprintf(out, "\n%s", scanner.Text())
	}
}

func (m Maven) Args() string {
	return strings.Join(m.cmd.Args, " ")
}

func (m *Maven) ModifyArgs(args []string) {
	m.cmd.Args = append([]string{"mvn", "--batch-mode"}, args...)
}
