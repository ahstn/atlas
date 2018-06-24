package config

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func validateConfig(s string) (string, error) {
	if !strings.Contains(s, ".yaml") {
		return "", errors.New("config should be a .yaml file")
	}

	if _, err := os.Stat(s); os.IsNotExist(err) {
		pwd := path.Join(os.Getenv("PWD"), filepath.Base(s))
		if _, err := os.Stat(pwd); os.IsNotExist(err) {
			return "", errors.New("config does not exist in ~/.config/atlas/ or pwd")
		}

		s = pwd
	}

	return s, nil
}

func validateRoot(s string) (string, error) {
	if strings.Contains(s, "~") {
		s = strings.Replace(s, "~", os.Getenv("HOME"), 1)
	}

	if _, err := os.Stat(s); os.IsNotExist(err) {
		return "", err
	}

	return s, nil
}
