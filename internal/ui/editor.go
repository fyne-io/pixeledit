package ui

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"path/filepath"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/widget"
	"golang.org/x/image/draw"

	"github.com/fyne-io/pixeledit/internal/api"
)

type editor struct {
	drawSurface             *interactiveRaster
	status                  *widget.Label
	cache                   *image.RGBA
	cacheWidth, cacheHeight int
	fgPreview               *canvas.Rectangle

	uri  string
	img  *image.RGBA
	zoom int
	fg   color.Color
	tool api.Tool

	win        fyne.Window
}

func (e *editor) PixelColor(x, y int) color.Color {
	return e.img.At(x, y)
}

func colorToBytes(col color.Color) []uint8 {
	r, g, b, a := col.RGBA()
	return []uint8{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

func (e *editor) SetPixelColor(x, y int, col color.Color) {
	i := (y*e.img.Bounds().Dx() + x) * 4
	rgba := colorToBytes(col)
	e.img.Pix[i] = rgba[0]
	e.img.Pix[i+1] = rgba[1]
	e.img.Pix[i+2] = rgba[2]
	e.img.Pix[i+3] = rgba[3]

	e.renderCache()
}

func (e *editor) FGColor() color.Color {
	return e.fg
}

func (e *editor) SetFGColor(col color.Color) {
	e.fg = col
	e.fgPreview.FillColor = col
	e.fgPreview.Refresh()
}

func (e *editor) buildUI() fyne.CanvasObject {
	return widget.NewScrollContainer(e.drawSurface)
}

func (e *editor) setZoom(zoom int) {
	e.zoom = zoom
	e.updateSizes()
	e.drawSurface.Refresh()
}

func (e *editor) setTool(tool api.Tool) {
	e.tool = tool
}

func (e *editor) draw(w, h int) image.Image {
	if e.cacheWidth == 0 || e.cacheHeight == 0 {
		return image.NewRGBA(image.Rect(0, 0, w, h))
	}

	if w > e.cacheWidth || h > e.cacheHeight {
		bigger := image.NewRGBA(image.Rect(0, 0, w, h))
		draw.Draw(bigger, e.cache.Bounds(), e.cache, image.Point{}, draw.Over)
		return bigger
	}

	return e.cache
}

func (e *editor) updateSizes() {
	if e.img == nil {
		return
	}
	e.cacheWidth = e.img.Bounds().Dx() * e.zoom
	e.cacheHeight = e.img.Bounds().Dy() * e.zoom

	c := fyne.CurrentApp().Driver().CanvasForObject(e.status)
	scale := float32(1.0)
	if c != nil {
		scale = c.Scale()
	}
	e.drawSurface.SetMinSize(fyne.NewSize(
		int(float32(e.cacheWidth)/scale),
		int(float32(e.cacheHeight)/scale)))

	e.renderCache()
}

func (e *editor) pixAt(x, y int) []uint8 {
	ix := x / e.zoom
	iy := y / e.zoom

	if ix >= e.img.Bounds().Dx() || iy >= e.img.Bounds().Dy() {
		return []uint8{0, 0, 0, 0}
	}

	return colorToBytes(e.img.At(ix, iy))
}

func (e *editor) renderCache() {
	e.cache = image.NewRGBA(image.Rect(0, 0, e.cacheWidth, e.cacheHeight))
	for y := 0; y < e.cacheHeight; y++ {
		for x := 0; x < e.cacheWidth; x++ {
			i := (y*e.cacheWidth + x) * 4
			col := e.pixAt(x, y)
			e.cache.Pix[i] = col[0]
			e.cache.Pix[i+1] = col[1]
			e.cache.Pix[i+2] = col[2]
			e.cache.Pix[i+3] = col[3]
		}
	}

	e.drawSurface.Refresh()
}

func fixEncoding(img image.Image) *image.RGBA {
	if rgba, ok := img.(*image.RGBA); ok {
		return rgba
	}

	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Over)
	return newImg
}

func (e *editor) LoadFile(read fyne.URIReadCloser) {
	defer read.Close()
	img, _, err := image.Decode(read)
	if err != nil {
		fyne.LogError("Unable to decode image", err)
		e.status.SetText(err.Error())
		return
	}

	e.uri = read.URI().String()
	e.img = fixEncoding(img)
	e.status.SetText(fmt.Sprintf("File: %s | Width: %d | Height: %d",
		filepath.Base(read.URI().String()), e.img.Bounds().Dx(), e.img.Bounds().Dy()))
	e.updateSizes()
}

func (e *editor) Reload() {
	if e.uri == "" {
		return
	}

	read, err := storage.OpenFileFromURI(storage.NewURI(e.uri))
	if err != nil {
		fyne.LogError("Unable to open file \""+e.uri+"\"", err)
		return
	}
	e.LoadFile(read)
}

func (e *editor) Save() {
	if e.uri == "" {
		return
	}

	uri := storage.NewURI(e.uri)
	if !e.isSupported(uri.Extension()) {
		fyne.LogError("Save only supports PNG", nil)
		return
	}
	write, err := storage.SaveFileToURI(uri)
	if err != nil {
		fyne.LogError("Error opening file to replace", err)
		return
	}

	e.saveToWriter(write)
}

func (e *editor) saveToWriter(write fyne.URIWriteCloser) {
	defer write.Close()
	if e.isPNG(write.URI().Extension()) {
		err := png.Encode(write, e.img)

		if err != nil {
			fyne.LogError("Could not encode image", err)
		}
	}
}

func (e *editor) SaveAs(writer fyne.URIWriteCloser) {
	e.saveToWriter(writer)
}

func (e *editor) isSupported(path string) bool {
	return e.isPNG(path)
}

func (e *editor) isPNG(path string) bool {
	return strings.LastIndex(strings.ToLower(path), "png") == len(path)-3
}

// NewEditor creates a new pixel editor that is ready to have a file loaded
func NewEditor() api.Editor {
	fgCol := color.Black
	edit := &editor{zoom: 1, fg: fgCol, fgPreview: canvas.NewRectangle(fgCol), status: newStatusBar()}
	edit.drawSurface = newInteractiveRaster(edit)

	return edit
}
