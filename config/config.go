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
	Docker Docker   `yaml:"docker"`
	Name   string   `yaml:"name"`
	Repo   string   `yaml:"repo"`
	Tasks  []string `yaml:"tasks"`
	Test   bool     `yaml:"test"`
}

// Docker stores container information relating to the build
type Docker struct {
	Dockerfile string `yaml:"dockerfile"`
	Enabled    bool   `yaml:"enabled"`
}

// Read attempts to parse the config file located at the path passed in
func Read(path string) (Project, error) {
	var p Project

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return Project{}, err
	}

	err = yaml.Unmarshal(file, &p)
	if err != nil {
		return Project{}, err
	}

	return p, nil
}
