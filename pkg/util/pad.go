package util

import (
	"math/rand"

	"github.com/fatih/color"
)

// PadLeft fills out the left of the string with spaces
func PadLeft(s string) (out string) {
	p := 12 - len(s)
	for i := 0; i < p; i++ {
		out += " "
	}
	return out + s
}

// PadLeftColor does the same as PadLeft but adds color to the string
func PadLeftColor(out func(...interface{}) string, s string) string {
	return out(PadLeft(s))
}

// RandomOutputColor uses fatih/color to return a function that will be used
// to output a string that is bold and colored
func RandomOutputColor() func(...interface{}) string {
	switch rand.Intn(6) {
	case 1:
		return color.New(color.FgGreen, color.Bold).SprintFunc()
	case 2:
		return color.New(color.FgYellow, color.Bold).SprintFunc()
	case 3:
		return color.New(color.FgBlue, color.Bold).SprintFunc()
	case 4:
		return color.New(color.FgMagenta, color.Bold).SprintFunc()
	case 5:
		return color.New(color.FgCyan, color.Bold).SprintFunc()
	default:
		return color.New(color.FgRed, color.Bold).SprintFunc()
	}
}
