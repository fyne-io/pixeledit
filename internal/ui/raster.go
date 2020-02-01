package ui

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type interactiveRaster struct {
	widget.BaseWidget

	edit *editor
	min  fyne.Size
	img  *canvas.Raster
}

func (r *interactiveRaster) SetMinSize(size fyne.Size) {
	r.min = size
	r.Resize(size)
}

func (r *interactiveRaster) MinSize() fyne.Size {
	return r.min
}

func (r *interactiveRaster) CreateRenderer() fyne.WidgetRenderer {
	return &rasterWidgetRender{raster: r, bg: canvas.NewRasterWithPixels(bgPattern)}
}

func (r *interactiveRaster) Tapped(ev *fyne.PointEvent) {
	if r.edit.tool == nil {
		return
	}

	x, y := r.locationForPosition(ev.Position)
	if x >= r.edit.img.Bounds().Dx() || y >= r.edit.img.Bounds().Dy() {
		return
	}

	r.edit.tool.Clicked(x, y, r.edit)
}

func (r *interactiveRaster) TappedSecondary(*fyne.PointEvent) {
}

func (r *interactiveRaster) locationForPosition(pos fyne.Position) (int, int) {
	c := fyne.CurrentApp().Driver().CanvasForObject(r.img)
	scale := float32(1.0)
	if c != nil {
		scale = c.Scale()
	}

	return int(float32(pos.X)*scale) / r.edit.zoom, int(float32(pos.Y)*scale) / r.edit.zoom
}

func newInteractiveRaster(edit *editor) *interactiveRaster {
	r := &interactiveRaster{img: canvas.NewRaster(edit.draw), edit: edit}
	r.ExtendBaseWidget(r)
	return r
}

type rasterWidgetRender struct {
	raster *interactiveRaster
	bg     *canvas.Raster
}

func bgPattern(x, y, _, _ int) color.Color {
	const boxSize = 25

	if (x/boxSize)%2 == (y/boxSize)%2 {
		return color.Gray{Y: 58}
	}

	return color.Gray{Y: 84}
}

func (r *rasterWidgetRender) Layout(size fyne.Size) {
	r.bg.Resize(size)
	r.raster.img.Resize(size)
}

func (r *rasterWidgetRender) MinSize() fyne.Size {
	return r.MinSize()
}

func (r *rasterWidgetRender) Refresh() {
	canvas.Refresh(r.raster)
}

func (r *rasterWidgetRender) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *rasterWidgetRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.raster.img}
}

func (r *rasterWidgetRender) Destroy() {
}
