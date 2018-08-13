package pb

import (
	"fmt"
	"time"

	spinner "github.com/briandowns/spinner"
)

// CreateAndStartBuildSpinner displays and returns a new spinner
func CreateAndStartBuildSpinner(title string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	_ = s.Color("fgHiCyan")
	s.Prefix = " "
	s.Suffix = fmt.Sprintf("  %s Building ...", title)
	s.FinalMSG = fmt.Sprintf("  ✔ %s Complete\n", title)
	s.Start()
	return s
}

// RunSpinner displays a new spinner for a set amount of time
func RunSpinner(title string) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	_ = s.Color("fgHiCyan")
	s.Prefix = " "
	s.Suffix = fmt.Sprintf(" %s...", title)
	s.FinalMSG = " ✔ Complete\n"
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()
}
