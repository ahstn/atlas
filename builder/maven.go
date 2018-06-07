package builder

import (
	"bufio"
	"bytes"
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

// Run executes the built command
func (m *Maven) Run() error {
	var stderr bytes.Buffer
	stdoutPipe, err := m.cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
	}

	stderrPipe, err := m.cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StderrPipe for Cmd", err)
	}

	queue := make([]*spinner.Spinner, 0)
	scanner := bufio.NewScanner(stdoutPipe)
	go func() {
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "Building") {
				if len(queue) != 0 {
					x := queue[0]
					x.Stop()
					queue = queue[1:]
				}

				spinner := pb.CreateAndStartBuildSpinner(scanner.Text()[20:])
				queue = append(queue, spinner)
			}
		}
	}()

	errScanner := bufio.NewScanner(stderrPipe)
	go func() {
		for errScanner.Scan() {
			fmt.Printf("[ERROR]: %s\n", errScanner.Text())
		}
	}()

	err = m.cmd.Start()
	if err != nil {
		fmt.Println("ERRROR")
		return err
	}

	err = m.cmd.Wait()
	if err != nil {
		fmt.Println("WAAAAIT")
		m.cmd.Stderr = &stderr
		fmt.Println(string(stderr.Bytes()))
		fmt.Println("HI", err)
		fmt.Println("Scanner:", string(errScanner.Bytes()))
		return err
	}

	return nil
}
