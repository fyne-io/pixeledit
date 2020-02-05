package tool

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type testEditor struct {
	img *image.RGBA

	fg color.Color
}

func (t *testEditor) BuildUI() fyne.CanvasObject {
	return widget.NewLabel("Not used")
}

func (t *testEditor) LoadFile(string) {
	// no-op
}

func (t *testEditor) Reload() {
	t.img = testImage()
}

func (t *testEditor) Save() {
	//no-op
}

func (t *testEditor) PixelColor(x, y int) color.Color {
	return t.img.At(x, y)
}

// TODO move out
func colorToBytes(col color.Color) []uint8 {
	r, g, b, a := col.RGBA()
	return []uint8{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

func (t *testEditor) SetPixelColor(x, y int, col color.Color) {
	i := (y*t.img.Bounds().Dx() + x) * 4
	rgba := colorToBytes(col)
	t.img.Pix[i] = rgba[0]
	t.img.Pix[i+1] = rgba[1]
	t.img.Pix[i+2] = rgba[2]
	t.img.Pix[i+3] = rgba[3]
}

func (t *testEditor) FGColor() color.Color {
	return t.fg
}

func (t *testEditor) SetFGColor(col color.Color) {
	t.fg = col
}

func testImage() *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, 4, 4))
}

func newTestEditor() *testEditor {
	return &testEditor{
		img: testImage(),
		fg:  color.Black,
	}
}
