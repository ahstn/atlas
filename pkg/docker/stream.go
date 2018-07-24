package docker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
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
	if s.ErrorMsg() != "" {
		return errors.New(s.ErrorMsg())
	}

	s.Stream = strings.TrimSpace(s.Stream)
	if strings.Contains(s.Stream, "Step") {
		fmt.Println(s.Stream)
	}

	return nil
}

// PrintStream decodes the Docker output from io.Reader and outputs it to
// the io.Writer
func PrintStream(r io.Reader) error {
	decoder := json.NewDecoder(r)

	var ds Stream
	for {
		if err := decoder.Decode(&ds); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		err := ds.Print()
		if err != nil {
			return err
		}
	}

	return nil
}
