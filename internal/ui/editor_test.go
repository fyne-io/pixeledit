package ui

import (
	"image/color"
	"path/filepath"
	"testing"

	_ "fyne.io/fyne/test" // load a test application

	"github.com/magiconair/properties/assert"
)

func testFile(name string) string {
	return filepath.Join(".", "testdata", name+".png")
}

func TestEditor_LoadFile(t *testing.T) {
	file := testFile("8x8")
	e := NewEditor()
	e.LoadFile(file)

	assert.Equal(t, color.RGBA{A: 255}, e.PixelColor(0, 0))
	assert.Equal(t, color.RGBA{R: 255, G: 255, B: 255, A: 255}, e.PixelColor(1, 0))
}

func TestEditor_SetPixelColor(t *testing.T) {
	file := testFile("8x8")
	e := NewEditor()
	e.LoadFile(file)

	assert.Equal(t, color.RGBA{A: 255}, e.PixelColor(0, 0))
	col := color.RGBA{R: 255, G: 255, B: 0, A: 128}
	e.SetPixelColor(1, 1, col)
	assert.Equal(t, col, e.PixelColor(1, 1))
}
