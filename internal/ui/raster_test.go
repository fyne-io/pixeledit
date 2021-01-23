package ui

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

func TestInteractiveRaster_MinSize(t *testing.T) {
	file := testFile("8x8")
	e := testEditorWithFile(file).(*editor)

	rast := newInteractiveRaster(e)
	e.drawSurface = rast
	e.setZoom(1)
	assert.Equal(t, fyne.NewSize(8, 8), rast.MinSize())
	assert.Equal(t, fyne.NewSize(8, 8), rast.Size())

	e.setZoom(2)
	assert.Equal(t, fyne.NewSize(16, 16), rast.MinSize())
	assert.Equal(t, fyne.NewSize(16, 16), rast.Size())
}

func TestInteractiveRaster_locationForPositon(t *testing.T) {
	file := testFile("8x8")
	e := testEditorWithFile(file).(*editor)

	r := newInteractiveRaster(e)
	x, y := r.locationForPosition(fyne.NewPos(2, 2))
	assert.Equal(t, 2, x)
	assert.Equal(t, 2, y)

	testCanvas := fyne.CurrentApp().Driver().CanvasForObject(r).(test.WindowlessCanvas)
	testCanvas.SetScale(2.0)
	x, y = r.locationForPosition(fyne.NewPos(2, 2))
	assert.Equal(t, 4, x)
	assert.Equal(t, 4, y)
}
