package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPadLeft(t *testing.T) {
	assert.Equal(t, PadLeft("test"), "        test")
	assert.Equal(t, PadLeftColor(RandomOutputColor(), "test"), "        test")
}
