package analysis

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/karrick/godirwalk"
)

const (
	Go         Language = "go"
	Java       Language = "java"
	JavaScript Language = "javascript"
	Unknown    Language = "unknown"
)

type (
	// Language is an enum of the supported programming languages.
	// Mainly used for making code more readable and comparisons easier
	Language string

	// Percentages is a map containing a repo's languages and the percentages
	// detailing the repo's makeup for each language
	Percentages map[Language]float64

	languageLines map[Language]int

	// RepoResult is the final result of processing a repo
	RepoResult struct {
		// Map containing all the languages detected and their overall percentage
		// makeup of the repo.
		Percentages Percentages

		// The primary language used in the repo, ultimately to save mutliple
		// map iterations across this codebase to determine the primary language.
		Language Language
	}
)

// DetectLanguage traverses the current directory and calculates the
// amount of language files they are (percentage based)
func DetectLanguage(dir string) (RepoResult, error) {
	languages := make(map[Language]int)
	var total int

	err := godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.ModeType() != os.ModeDir {
				// We only want to count valid languages
				if ext := getExtension(osPathname); ext != Unknown {
					lines, _ := lineCounter(osPathname)
					languages[getExtension(osPathname)] += lines
					total += lines
					fmt.Printf("%s : %d (%d)\n", getExtension(osPathname), lines, total)
				}
			}
			return nil
		},
		Unsorted: true,
	})

	if err != nil {
		return RepoResult{}, err
	}

	// Calculate percentages at the end rather than after every file
	// Might be better keeping an ordered slice (highest at top)
	return calculatePercentages(languages, total)
}

func getExtension(filename string) Language {
	f := strings.Split(filename, ".")

	switch f[len(f)-1] {
	case string(Go):
		return Go
	case string(Java):
		return Java
	case string(JavaScript), "js", "jsx", "ts":
		return JavaScript
	default:
		return Unknown
	}
}

func lineCounter(filename string) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	file, _ := os.Open(filename)
	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func calculatePercentages(languages map[Language]int, total int) (RepoResult, error) {
	percentages := make(Percentages)
	primaryLanguage := Unknown
	primaryPercentage := 0.0

	for k, v := range languages {
		percentage := float64(v) / float64(total) * 100
		percentages[k] = percentage

		if percentage > 100 {
			return RepoResult{}, errors.New("error calculating repo percentages")
		}
		if percentage > primaryPercentage {
			primaryLanguage = k
		}

		fmt.Printf("(%d / %d) * 100 = %f\n", v, total, float32(v)/float32(total)*100)
	}

	return RepoResult{
		Percentages: percentages,
		Language:    primaryLanguage,
	}, nil
}
