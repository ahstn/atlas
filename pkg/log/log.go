package log

import (
	"fmt"
	"io"
	"os"
	"strings"

	emoji "gopkg.in/kyokomi/emoji.v1"
)

//go:generate mockery -name Logger -case underscore

// Logger defines the actions that a CLI logger should preform
type Logger interface {
	Print(string)
	Printf(string, ...interface{})
	CheckError(error)
}

// Client is the implementation for Logger
type Client struct {
	out io.Writer
}

// Print displays the passed string in the client's out writer.
// If the string begins with ':' and two or more ':' characters,
// it should be a safe assumption that it contains an emoji shortcode.
// i.e. :angel: or :whale:
func (c Client) Print(s string) {
	if strings.TrimSpace(s)[0] == ':' && strings.Count(s, ":") > 1 {
		s = emoji.Sprint(s)
	}

	fmt.Fprint(c.out, s)
}

// Printf displays the passed string in the client's out writer using the
// format and args specified.
func (c Client) Printf(s string, args ...interface{}) {
	if strings.TrimSpace(s)[0] == ':' && strings.Count(s, ":") > 1 {
		s = emoji.Sprintf(s, args)
	}

	fmt.Fprintf(c.out, s, args)
}

// CheckError provides error checking and prints a stack if the debug env var
// is set
func (c Client) CheckError(err error) {
	if err == nil {
		return
	}

	emoji.Fprintln(c.out, ":x: Error occurred:", err.Error())
	if os.Getenv("ATLAS_DEBUG") == "1" {
		panic(err)
	}

	if os.Getenv("ATLAS_UNIT_TEST") != "1" {
		os.Exit(1)
	}
}

// NewClient returns a new logger instance with out set to Stdout
func NewClient() Client {
	return Client{
		out: os.Stdout,
	}
}
