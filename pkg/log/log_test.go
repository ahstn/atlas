package log

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	c := Client{out: &buf}

	c.Print("Hello")
	assert.Contains(t, buf.String(), "Hello")

	c.Print(":whale: Docker!")
	assert.Contains(t, buf.String(), "\U0001f433") // Emoji Unicode
	assert.Contains(t, buf.String(), "Docker!")
}

func TestPrintf(t *testing.T) {
	var buf bytes.Buffer
	c := Client{out: &buf}

	c.Printf(":whale: %s", "Docker!")
	assert.Contains(t, buf.String(), "\U0001f433") // Emoji Unicode
	assert.Contains(t, buf.String(), "Docker!")
}

func TestCheckError(t *testing.T) {
	var buf bytes.Buffer
	c := Client{out: &buf}

	c.CheckError(nil)
	assert.Empty(t, buf.String())

	os.Setenv("ATLAS_UNIT_TEST", "1")
	c.CheckError(io.EOF)
	assert.Contains(t, buf.String(), "\u274c") // Emoji Unicode
	assert.Contains(t, buf.String(), "EOF")
}

func TestCheckErrorPanics(t *testing.T) {
	var buf bytes.Buffer
	c := Client{out: &buf}

	os.Setenv("ATLAS_DEBUG", "1")
	panics := false
	var msg interface{}
	func() {
		defer func() {
			if msg = recover(); msg != nil {
				panics = true
			}
		}()
		c.CheckError(io.EOF)
	}()

	assert.Equal(t, true, panics)
}
