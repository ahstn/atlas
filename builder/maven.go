package builder

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ahstn/atlas/pb"
	"github.com/apex/log"
	"github.com/briandowns/spinner"
)

func getPath() (string, error) {
	return "/usr/bin/mvn", nil
}

// Maven is a implmentation of Builder{} for Java
type Maven struct {
	cmd exec.Cmd
}

func (m *Maven) initialiseCommand() {
	if m.cmd.Path == "" {
		path, err := getPath()
		if err != nil {
			log.Info("unable to find path (mvn)")
		}
		m.cmd = exec.Cmd{
			Path: path, // allow user to set custom exec path
			Args: []string{""},
			Env:  nil, // allow user to set custom environment variables
			Dir:  "",  // allow user to pass custom project path
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
	if v {
		go printVerboseLog(scanner)
	} else {
		go printLog(scanner) // TODO: WaitGroup here
	}

	errScanner := bufio.NewScanner(stderrPipe)
	go func() {
		for errScanner.Scan() {
			//fmt.Printf("[ERROR]: %s\n", errScanner.Text())
		}
	}()

	err = m.cmd.Start()
	if err != nil {
		fmt.Println("ERRROR")
		return err
	}

	err = m.cmd.Wait()
	if err != nil {
		fmt.Println("\n ")
		return err
	}

	return nil
}

func printLog(scanner *bufio.Scanner) {
	failedTests := false
	queue := make([]*spinner.Spinner, 0)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Failed tests:") {
			failedTests = true
		} else if strings.Contains(scanner.Text(), "Tests run:") {
			failedTests = false
		}

		if failedTests {
			fmt.Printf("\n%s", scanner.Text())
		} else if strings.Contains(scanner.Text(), "Failed to execute goal") {
			fmt.Printf("\n\n%s\n", scanner.Text())
		} else if strings.Contains(scanner.Text(), "Building") {
			// If another "Building" string is detected, last build is done
			// therefore update the last spinner
			if len(queue) != 0 {
				x := queue[0]
				x.Stop()
				queue = queue[1:]
			}

			// Create spinner and add it to the queue of pending builds
			spinner := pb.CreateAndStartBuildSpinner(scanner.Text()[20:])
			queue = append(queue, spinner)
		}
	}

	queue[len(queue)-1].Stop()
}

func printVerboseLog(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Printf("\n%s", scanner.Text())
	}
}
