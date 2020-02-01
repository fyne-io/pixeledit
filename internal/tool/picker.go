package tool

import (
	"fyne.io/fyne"

	"github.com/fyne-io/pixeledit/internal/api"
	"github.com/fyne-io/pixeledit/internal/data"
)

// Picker allows setting the current image colors from a pixel in an image
type Picker struct {
}

// Name returns the name of this tool
func (p *Picker) Name() string {
	return "Picker"
}

// Icon returns the icon for this tool
func (p *Picker) Icon() fyne.Resource {
	return data.DropperIcon
}

// Clicked will set the editor foreground color to the color of the pixel under the cursor
func (p *Picker) Clicked(x, y int, edit api.Editor) {
	edit.SetFGColor(edit.PixelColor(x, y))
}
