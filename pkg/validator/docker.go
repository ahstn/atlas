package validator

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// ValidateArguments verifies that the path argument is present
// If the path is pointing to a file, it returns the directory instead
// TODO: Default to pwd?
func ValidateArguments(s string) (string, error) {
	if s == "" {
		return "", errors.New("invalid argument - you must pass a directory")
	}

	info, err := os.Stat(s)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return filepath.Abs(path.Dir(s))
	}

	return filepath.Abs(s)
}

// ValidateTag verifies that the tag argument is valid
func ValidateTag(s string) error {
	if !strings.Contains(s, ":") {
		return errors.New("invalid tag - you must specify a version")
	} else if s[0] == '/' || s[0] == ':' {
		return errors.New("invalid tag - must not start with '/' or ':'")
	} else if s[len(s)-1] == '/' || s[len(s)-1] == ':' {
		return errors.New("invalid tag - must not end with '/' or ':'")
	}

	return nil
}

// ValidateBuildArgs verifies that the args present are valid
func ValidateBuildArgs(s []string) error {
	for _, arg := range s {
		fmt.Println(arg)
		if !strings.Contains(arg, "=") {
			return errors.New("arg has no value - must be in `key=value` format")
		} else if arg[len(arg)-1] == '=' {
			return errors.New("arg has no value - must be in `key=value` format")
		} else if strings.Count(arg, "=") > 1 {
			return errors.New("invalid arg - must be in `key=value` format")
		}
	}

	return nil
}

// TryFindDockerfile scans a base path for any file matching 'Dockerfile'
func TryFindDockerfile(base string) (string, error) {
	var df string
	filepath.Walk(base, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString("Dockerfile", f.Name())
			if err == nil && r {
				df = f.Name()
				return nil
			}
		}

		return nil
	})

	if df == "" {
		return "", errors.New("no Dockerfile specifed and unable to locate one")
	}

	return df, nil
}
