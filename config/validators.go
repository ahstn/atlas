package config

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// ValidateExists verifies that the config file is present
// First checks the CWD, then ~/.config/atlas
// param: s should the filename (not filepath)
func ValidateExists(s string) (string, error) {
	pwd := path.Join(os.Getenv("PWD"), filepath.Base(s))
	home := path.Join(os.Getenv("HOME"), ".config.atlas/", filepath.Base(s))

	if _, err := os.Stat(pwd); os.IsNotExist(err) {
		if _, err := os.Stat(home); os.IsNotExist(err) {
			return "", errors.New("config does not exist in ~/.config/atlas/ or pwd")
		}
		return home, nil
	}
	return pwd, nil
}

// ValidateConfig verifies that the filetype and contents are valid
// param: s should be the full fil path
func ValidateConfig(s string) error {
	if !strings.Contains(s, ".yaml") {
		return errors.New("config should be a .yaml file")
	}

	return nil
}

// ValidateConfigBaseDir ensures the base dir in the config actually exists
// and converts it to a full path rather than using '~' shorthand
func ValidateConfigBaseDir(s string) (string, error) {
	if strings.Contains(s, "~") {
		s = strings.Replace(s, "~", os.Getenv("HOME"), 1)
	}

	if _, err := os.Stat(s); os.IsNotExist(err) {
		return "", err
	}

	return s, nil
}
