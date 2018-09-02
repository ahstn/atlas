package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/ahstn/atlas/pkg/validator"
	yaml "gopkg.in/yaml.v2"
)

// Project is what define an collection of applications to be built
type Project struct {
	Root     string     `yaml:"root"`
	Services []*Service `yaml:"services"`
}

// Service is a single buildable application
type Service struct {
	Docker  DockerArtifact `yaml:"docker" json:"docker,omitempty"`
	Package Package        `yaml:"package" json:"package,omitempty"`
	Name    string         `yaml:"name" json:"name,omitempty"`
	Repo    string         `yaml:"repo" json:"repo,omitempty"`
	Path    string         `yaml:"path" json:"path,omitempty"`
	Tasks   []string       `yaml:"tasks" json:"tasks,omitempty"`
	Args    []string       `yaml:"args" json:"args,omitempty"`
	Env     []string       `yaml:"env" json:"env,omitempty"`
	Test    bool           `yaml:"test" json:"test,omitempty"`
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
	Env        []string `yaml:"env"`
	Ports      []string `yaml:"port"`
	Dockerfile string   `yaml:"dockerfile"`
	Path       string   `yaml:"path"`
	Enabled    bool     `yaml:"enabled"`
	Tag        string   `yaml:"tag"`
	Cmd        string   `yaml:"cmd"`
}

// Package stores packaging information relating to the build
type Package struct {
	Args   string `yaml:"args"`
	SubDir string `yaml:"subDir"`
}

// Read attempts to parse the config file located at the path passed in
// TODO: Seperate this into two functions to make testing easier
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

	s, err := validator.ValidateConfigBaseDir(p.Root)
	if err != nil {
		return Project{}, err
	}
	p.Root = s

	// Validate Docker fields and add metadata
	for _, svc := range p.Services {
		svc.Path = filepath.Join(p.Root, svc.Name)
		svc.Docker.Path = filepath.Join(svc.Path, svc.Docker.Path)

		if err := validator.ValidateBuildArgs(svc.Docker.Args); err != nil {
			return Project{}, err
		}

		if svc.Docker.Dockerfile == "" {
			svc.Docker.Dockerfile, err = validator.TryFindDockerfile(svc.Path)
			if err != nil {
				return Project{}, err
			}
		}

		svc.Docker.Dockerfile = filepath.Join(svc.Path, svc.Docker.Dockerfile)
	}

	return p, nil
}
