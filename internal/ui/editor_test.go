package ui

import (
	"image/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	_ "fyne.io/fyne/test" // load a test application

	"github.com/stretchr/testify/assert"
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

func TestEditor_Reset(t *testing.T) {
	file := testFile("8x8")
	e := NewEditor()
	e.LoadFile(file)

	assert.Equal(t, color.RGBA{A: 255}, e.PixelColor(0, 0))

	red := color.RGBA{255, 0, 0, 255}
	e.SetPixelColor(0, 0, red)
	assert.Equal(t, red, e.PixelColor(0, 0))

	e.Reload()
	assert.Equal(t, color.RGBA{A: 255}, e.PixelColor(0, 0))
}

func TestEditor_Save(t *testing.T) {
	origFile := testFile("8x8")
	file := testFile("8x8-tmp")
	content, err := ioutil.ReadFile(origFile)
	if err != nil {
		t.Error("Failed to read test file")
	}
	err = ioutil.WriteFile(file, content, 0644)
	if err != nil {
		t.Error("Failed to write temporary file")
	}
	defer func() {
		err = os.Remove(file)
		if err != nil {
			t.Error("Failed to remove temporary file")
		}
	}()

	e := NewEditor()
	e.LoadFile(file)

	assert.Equal(t, color.RGBA{A: 255}, e.PixelColor(0, 0))

	red := color.RGBA{255, 0, 0, 255}
	e.SetPixelColor(0, 0, red)
	assert.Equal(t, red, e.PixelColor(0, 0))

	e.Save()

	e.LoadFile(file)
	assert.Equal(t, red, e.PixelColor(0, 0))
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
