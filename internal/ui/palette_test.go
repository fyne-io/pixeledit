package ui

import (
	"testing"

	"fyne.io/fyne"
	"fyne.io/fyne/test"
	"fyne.io/fyne/widget"
	"github.com/stretchr/testify/assert"
)

func TestDefaultZoom(t *testing.T) {
	file := testFile("8x8")
	e := testEditorWithFile(file)

	p := newPalette(e.(*editor))
	zoom := p.(*widget.Box).Children[0].(*fyne.Container).Objects[1].(*widget.Box).Children[1].(*widget.Label)
	assert.Equal(t, "100%", zoom.Text)
}

func TestZoomIn(t *testing.T) {
	file := testFile("8x8")
	e := testEditorWithFile(file)
	assert.Equal(t, 1, e.(*editor).zoom)

	p := newPalette(e.(*editor))
	zoomItems := p.(*widget.Box).Children[0].(*fyne.Container).Objects[1].(*widget.Box).Children
	zoom := zoomItems[1].(*widget.Label)
	zoomIn := zoomItems[2].(*widget.Button)
	test.Tap(zoomIn)

	assert.Equal(t, 2, e.(*editor).zoom)
	assert.Equal(t, "200%", zoom.Text)
}
