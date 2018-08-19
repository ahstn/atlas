package pb

import (
	"fmt"
	"io"
	"time"

	spinner "github.com/briandowns/spinner"
)

// Print displays and returns a new spinner
func Print(title string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	_ = s.Color("fgHiCyan")
	s.Prefix = " "
	s.Suffix = fmt.Sprintf("  %s Building ...", title)
	s.FinalMSG = fmt.Sprintf("  ✔ %s Complete\n", title)
	s.Start()
	return s
}

// Fprint displays to the passed io.Writer and returns a new spinner
func Fprint(out io.Writer, title string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	_ = s.Color("fgHiCyan")
	s.Writer = out
	s.Prefix = " "
	s.Suffix = fmt.Sprintf("  %s Building ...", title)
	s.FinalMSG = fmt.Sprintf("  ✔ %s Complete\n", title)
	s.Start()
	return s
}
