package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Project is what define an collection of applications to be built
type Project struct {
	Root     string    `yaml:"root"`
	Services []Service `yaml:"services"`
}

// Service is a single buildable application
type Service struct {
	DockerArtifact DockerArtifact `yaml:"docker"`
	Package        Package        `yaml:"package"`
	Name           string         `yaml:"name"`
	Repo           string         `yaml:"repo"`
	Tasks          []string       `yaml:"tasks"`
	Test           bool           `yaml:"test"`
}

// HasTask is a helper function for detecting package task
func (s Service) HasTask(x string) bool {
	for _, t := range s.Tasks {
		if t == x {
			return true
		}
	}
	return false
}

// HasPackageSubDir is a helper function for detecting package subdir presence
func (s Service) HasPackageSubDir() bool {
	return s.Package.SubDir != ""
}

// DockerArtifact stores container information relating to the build
type DockerArtifact struct {
	Args       []string `yaml:"args"`
	Dockerfile string   `yaml:"dockerfile"`
	Path       string   `yaml:"path"`
	Enabled    bool     `yaml:"enabled"`
	Tag        string   `yaml:"tag"`
}

// Package stores packaging information relating to the build
type Package struct {
	Parameters string `yaml:"parameters"`
	SubDir     string `yaml:"subDir"`
}

// Read attempts to parse the config file located at the path passed in
func Read(path string) (Project, error) {
	var p Project
	var err error

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return Project{}, err
	}

	err = yaml.Unmarshal(file, &p)
	if err != nil {
		return Project{}, err
	}

	s, err := ValidateConfigBaseDir(p.Root)
	if err != nil {
		return Project{}, err
	}
	p.Root = s

	return p, nil
}
