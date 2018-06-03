package builder

import (
	"os/exec"

	"github.com/apex/log"
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
	err := m.cmd.Start()
	if err != nil {
		return err
	}

	err = m.cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
