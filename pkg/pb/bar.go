package pb

import (
	"fmt"
	"time"

	pb "gopkg.in/cheggaaa/pb.v2"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

// RunProgressBar prints a progress bar based on time
func RunProgressBar(title string) {
	tmpl := fmt.Sprintf(
		`{{"%s"}} {{bar . "|" "██" "░" "░" "|" | green}} {{speed . | blue }}`,
		emoji.Sprintf(" :wrench:%s...", title),
	)

	count := 1000
	bar := pb.ProgressBarTemplate(tmpl).Start(count)
	bar.SetWidth(80)
	for i := 0; i < count/2; i++ {
		bar.Add(2)
		time.Sleep(time.Millisecond * 4)
	}
	bar.Finish()
}
