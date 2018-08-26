package docker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ahstn/atlas/pkg/pb"
	"github.com/briandowns/spinner"
)

// StreamError defines an error that occured during a Docker event
type StreamError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Stream is used by Docker for communicating event responses
type Stream struct {
	Stream string      `json:"stream,omitempty"`
	Status string      `json:"status,omitempty"`
	Error  StreamError `json:"errorDetail,omitempty"`
}

// ErrorMsg is a helper method for fetching any errors
func (s Stream) ErrorMsg() string {
	return s.Error.Message
}

// Print outputs any valid stream content to the io.Writer passed in
func (s Stream) Print() error {

	s.Stream = strings.TrimSpace(s.Stream)
	if strings.Contains(s.Stream, "Step") {
		fmt.Println(s.Stream)

	}

	return nil
}

// PrintRun tails the logs from a running docker container with the app name
// prefixed to each line
func PrintRun(r io.Reader, app string) error {
	buf := make([]byte, 256)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf(" %s | %s\n", app, strings.TrimSpace(string(buf[:n])))
	}

	return nil
}

// PrintStream decodes the Docker output from io.Reader and outputs it to
// the io.Writer
func PrintStream(r io.Reader) error {
	decoder := json.NewDecoder(r)

	var ds Stream
	queue := make([]*spinner.Spinner, 0)
	for {
		if err := decoder.Decode(&ds); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if ds.ErrorMsg() != "" {
			return errors.New(ds.ErrorMsg())
		}

		if strings.Contains(ds.Stream, "Step") {
			formatted := strings.Replace(ds.Stream, "Step ", "[", 1)
			formatted = strings.Replace(formatted, " : ", "]: ", 1)

			if len(queue) != 0 {
				x := queue[0]
				x.Stop()
				queue = queue[1:]
			}

			spinner := pb.Print(formatted)
			queue = append(queue, spinner)
		}

	}
	for _, x := range queue {
		x.Stop()
		queue = queue[1:]
	}

	return nil
}
