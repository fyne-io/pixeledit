package api

import (
	"fyne.io/fyne"
)

// Tool represents any pixel editing tool that can be loaded into our editor
type Tool interface {
	Name() string                  // Name is the human readable name of this tool
	Icon() fyne.Resource           // Icon returns a resource that we can use for button icons
	Clicked(x, y int, edit Editor) // Clicked is called when the tool is active and the user interacts with the editor
}
