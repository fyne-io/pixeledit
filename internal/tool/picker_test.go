package tool

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPicker_Data(t *testing.T) {
	p := &Picker{}

	assert.Equal(t, "Picker", p.Name())
	assert.NotNil(t, p.Icon())
}

func TestPicker_Clicked(t *testing.T) {
	e := newTestEditor()
	col := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	e.SetPixelColor(0, 0, col)

	p := &Picker{}
	p.Clicked(0, 0, e)
	assert.Equal(t, col, e.FGColor())
}
