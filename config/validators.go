package config

import (
	"os"
	"strings"
)

func validateRoot(s string) (string, error) {
	if strings.Contains(s, "~") {
		s = strings.Replace(s, "~", os.Getenv("HOME"), 1)
	}

	if _, err := os.Stat(s); os.IsNotExist(err) {
		return "", err
	}

	return s, nil
}
