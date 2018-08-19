package pb

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestFprint(t *testing.T) {
	var buf bytes.Buffer

	spinner := Fprint(&buf, "spinner test")
	time.Sleep(time.Millisecond * 4)
	spinner.Stop()

	// t.Fatal(buf.String())
	if !strings.Contains(buf.String(), "⠋  spinner test Building ...") {
		t.Fatal("Expected Spinner to print 'Building ...'. Got:", buf.String())
	} else if !strings.Contains(buf.String(), "✔ spinner test Complete") {
		t.Fatal("Expected Spinner to print 'Complete'. Got:", buf.String())
	}
}
