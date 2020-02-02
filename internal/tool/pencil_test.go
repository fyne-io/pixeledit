package tool

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPencil_Data(t *testing.T) {
	p := &Pencil{}

	assert.Equal(t, "Pencil", p.Name())
	assert.NotNil(t, p.Icon())
}

func TestPencil_Clicked(t *testing.T) {
	e := newTestEditor()
	col := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	e.SetFGColor(col)

	p := &Pencil{}
	p.Clicked(0, 0, e)
	assert.Equal(t, col, e.PixelColor(0, 0))
}
