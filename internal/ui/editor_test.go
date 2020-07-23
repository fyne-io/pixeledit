package ui

import (
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"fyne.io/fyne"
	"fyne.io/fyne/storage"
	_ "fyne.io/fyne/test" // load a test application

	"github.com/stretchr/testify/assert"
)

func uriForTestFile(name string) fyne.URI {
	path := filepath.Join(".", "testdata", name+".png")

	return storage.NewURI("file://" + path)
}

func testFile(name string) fyne.URIReadCloser {
	read, err := storage.OpenFileFromURI(uriForTestFile(name))
	if err != nil {
		fyne.LogError("Unable to open file \""+name+"\"", err)
		return nil
	}

	return read
}

func testFileWrite(name string) fyne.URIWriteCloser {
	write, err := storage.SaveFileToURI(uriForTestFile(name))
	if err != nil {
		fyne.LogError("Unable to save file \""+name+"\"", err)
		return nil
	}

	return write
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
	outFile := testFileWrite("8x8-tmp")
	content, err := ioutil.ReadAll(origFile)
	if err != nil {
		t.Error("Failed to read test file")
	}
	_, err = outFile.Write([]byte(content))
	if err != nil {
		t.Error("Failed to write temporary file")
	}
	defer func() {
		err = os.Remove(outFile.URI().String()[7:])
		if err != nil {
			t.Error("Failed to remove temporary file")
		}
	}()

	e := NewEditor()
	file := testFile("8x8-tmp")
	e.LoadFile(file)

	assert.Equal(t, color.RGBA{A: 255}, e.PixelColor(0, 0))

	red := color.RGBA{255, 0, 0, 255}
	e.SetPixelColor(0, 0, red)
	assert.Equal(t, red, e.PixelColor(0, 0))

	e.Save()

	e.LoadFile(file)
	assert.Equal(t, red, e.PixelColor(0, 0))
}

func TestEditorFGColor(t *testing.T) {
	e := NewEditor()

	assert.Equal(t, color.Black, e.FGColor())
}

func TestEditor_SetFGColor(t *testing.T) {
	e := NewEditor()

	fg := color.White
	e.SetFGColor(fg)
	assert.Equal(t, fg, e.FGColor())
}

func TestEditor_PixelColor(t *testing.T) {
	file := testFile("8x8")
	e := NewEditor()
	e.LoadFile(file)

	assert.Equal(t, color.RGBA{A: 255}, e.PixelColor(0, 0))
	assert.Equal(t, color.RGBA{R: 0, G: 0, B: 0, A: 0}, e.PixelColor(9, 9))
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

func TestEditor_fixEncoding(t *testing.T) {
	size := 4
	nonRGBA := image.NewCMYK(image.Rect(0, 0, size, size))

	fixed := fixEncoding(nonRGBA)
	assert.Equal(t, image.Rect(0, 0, size, size), fixed.Bounds())
}

func TestEditor_isSupported(t *testing.T) {
	e := NewEditor().(*editor)

	assert.True(t, e.isSupported("test.png"))
	assert.False(t, e.isPNG("test.jpg"))
}

func TestEditor_isPNG(t *testing.T) {
	e := NewEditor().(*editor)

	assert.True(t, e.isPNG("test.png"))
	assert.True(t, e.isPNG("BIG.PNG"))
	assert.False(t, e.isPNG("wrong.ping"))
}
