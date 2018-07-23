package pb

import (
	"fmt"
	"time"

	spinner "github.com/briandowns/spinner"
)

func CreateAndStartBuildSpinner(title string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	s.Color("blue")
	s.Prefix = " "
	s.Suffix = fmt.Sprintf("  %s Building ...", title)
	s.FinalMSG = fmt.Sprintf("  ✔ %s Complete\n", title)
	s.Start()
	return s
}

func RunSpinner(title string) {
	s := spinner.New(spinner.CharSets[24], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	s.Color("blue")
	s.Prefix = " "
	s.Suffix = fmt.Sprintf(" %s...", title)
	s.FinalMSG = " ✔ Complete\n"
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()
}
