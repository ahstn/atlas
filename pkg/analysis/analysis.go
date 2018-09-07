package analysis

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/karrick/godirwalk"
)

type language string
type languageLines map[language]int

const (
	Go         language = "go"
	Java       language = "java"
	JavaScript language = "javascript"
	Unknown    language = "unknown"
)

// DetectLanguage traverses the current directory and calculates the
// amount of language files they are (percentage based)
func DetectLanguage(dir string) error {
	var languages map[language]int
	var total int

	err := godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.ModeType() != os.ModeDir {
				lines, _ := lineCounter(osPathname)
				languages[getExtension(osPathname)] += lines
				total++
			}
			fmt.Printf("%s %s\n", de.ModeType(), osPathname)
			return nil
		},
		Unsorted: true,
	})

	// Calculate percentages at the end rather than after every file
	// Might be better keeping an ordered slice (highest at top)
	percentages := calculatePercentages(languages, total)
	fmt.Println(percentages)

	return err
}

func getExtension(filename string) language {
	f := strings.Split(filename, ".")

	switch f[len(f)-1] {
	case string(Go):
		return Go
	case string(Java):
		return Java
	case string(JavaScript):
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

func calculatePercentages(languages map[language]int, total int) map[language]float32 {
	var percentages map[language]float32
	for k, v := range languages {
		percentages[k] = float32(v/total) * 100
	}

	return percentages
}
