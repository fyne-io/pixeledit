package ui

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	file := testFile("8x8")
	e := NewEditor()
	e.LoadFile(file)

	assert.True(t, strings.Contains(e.(*editor).status.Text, "File: 8x8.png"))
	assert.True(t, strings.Contains(e.(*editor).status.Text, "Width: 8"))
}
