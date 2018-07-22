package builder

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/ahstn/atlas/pb"
	"github.com/apex/log"
	"github.com/briandowns/spinner"
)

// Maven is a implmentation of Builder{} for Java
type Maven struct {
	Dir string
	cmd exec.Cmd
}

func (m *Maven) initialiseCommand() {
	if m.cmd.Path == "" {
		path, err := exec.LookPath("mvn")
		if err != nil {
			log.Info("unable to find path (mvn)")
		}
		m.cmd = exec.Cmd{
			Path: path, // allow user to set custom exec path
			Args: []string{""},
			Env:  nil, // allow user to set custom environment variables
			Dir:  m.Dir,
		}
	}
}

// Clean runs "mvn clean"
func (m *Maven) Clean() {
	m.initialiseCommand()
	m.cmd.Args = append(m.cmd.Args, "clean")
}

// Build runs "mvn build"
func (m *Maven) Build() {
	m.initialiseCommand()
	m.cmd.Args = append(m.cmd.Args, "install")
}

// Package runs "mvn package"
func (m *Maven) Package() {
	m.initialiseCommand()
	m.cmd.Args = append(m.cmd.Args, "package")
}

// SkipTests appends "-DskipTests" to the command
func (m *Maven) SkipTests() {
	m.initialiseCommand()
	m.cmd.Args = append(m.cmd.Args, "-DskipTests")
}

// SetDir changes the directory the command will execute in
func (m *Maven) SetDir(d string) {
	m.initialiseCommand()
	m.cmd.Dir = d
}

// Run executes the built command
func (m *Maven) Run(v bool) error {
	stdoutPipe, err := m.cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
	}

	stderrPipe, err := m.cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StderrPipe for Cmd", err)
	}

	scanner := bufio.NewScanner(stdoutPipe)
	var wg sync.WaitGroup
	if v {
		go printVerboseLog(scanner)
	} else {
		wg.Add(1)
		go printLog(scanner, &wg)
	}

	errScanner := bufio.NewScanner(stderrPipe)
	go func() {
		for errScanner.Scan() {
			//fmt.Printf("[ERROR]: %s\n", errScanner.Text())
		}
	}()

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

func printLog(s *bufio.Scanner, wg *sync.WaitGroup) {
	failedTests := false
	queue := make([]*spinner.Spinner, 0)
	//TODO: Handle Packaging output
	for s.Scan() {
		if strings.Contains(s.Text(), "Failed tests:") {
			failedTests = true
		} else if strings.Contains(s.Text(), "Tests run:") {
			failedTests = false
		}

		if failedTests {
			fmt.Printf("\n%s", s.Text())
		} else if strings.Contains(s.Text(), "Failed to execute goal") {
			fmt.Printf("\n\n%s\n", s.Text())
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
			spinner := pb.CreateAndStartBuildSpinner(module)
			queue = append(queue, spinner)
		}
	}

	for _, x := range queue {
		x.Stop()
		queue = queue[1:]
	}

	wg.Done()
}

func printVerboseLog(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Printf("\n%s", scanner.Text())
	}
}
