package api

import (
	"image/color"

	"fyne.io/fyne"
)

// Editor describes the editing capabilities of a pixel editor
type Editor interface {
	BuildUI(fyne.Window)         // BuildUI Loads the main editor GUI
	LoadFile(fyne.URIReadCloser) // LoadFile specifies a data stream to load from
	Reload()                     // Reload will reset the image to its original state
	Save()                       // Save writes the image back to its source location

	PixelColor(x, y int) color.Color         // Get the color of a pixel in our image
	SetPixelColor(x, y int, col color.Color) // Set the color of the indicated pixel

	FGColor() color.Color
	SetFGColor(color.Color)
}
