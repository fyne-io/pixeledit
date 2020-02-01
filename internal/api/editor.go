package api

import (
	"image/color"

	"fyne.io/fyne"
)

// Editor describes the editing capabilities of a pixel editor
type Editor interface {
	BuildUI() fyne.CanvasObject // BuildUI Loads the main editor GUI
	LoadFile(string)            // SetFile specifies a file to load from the filesystem

	PixelColor(x, y int) color.Color         // Get the color of a pixel in our image
	SetPixelColor(x, y int, col color.Color) // Set the color of the indicated pixel

	FGColor() color.Color
	SetFGColor(color.Color)
}
