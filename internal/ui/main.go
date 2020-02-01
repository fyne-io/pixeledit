package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
)

// BuildUI creates the main window of our pixel edit application
func (e *editor) BuildUI() fyne.CanvasObject {
	palette := newPalette(e)

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, e.status, palette, nil),
		e.status, palette, e.buildUI())
}
