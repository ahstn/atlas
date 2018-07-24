package util

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// OpenBrowser opens the default browser at the supplied URL
func OpenBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}

// ProcessRepoURL sanitizes URL to return Git repo page
func ProcessRepoURL(r string) (string, error) {
	if strings.Contains(r, "git@") {
		r = strings.Replace(r, "git@", "https://", 1)
		r = strings.Replace(r, ".com:", ".com/", 1)

		return r, nil
	} else if strings.Contains(r, "https://") {
		return r, nil
	}

	return "", errors.New("could not process Git repo url")
}

//TODO: Abstract logic from ProcessRepoURL + ProcessIssuesURL
// ProcessIssuesURL sanitizes URL
// func ProcessIssuesURL(r string) (string, error) {
// 	if strings.Contains(r, "git@") {
// 		r = strings.Replace(r, "git@", "https://", 1)
// 		r = strings.Replace(r, ".com:", ".com/", 1)

// 		return r, nil
// 	} else if strings.Contains(r, "https://") {
// 		return r, nil
// 	}

// 	return "", errors.New("could not process Git repo url")
// }
