package ui

import (
	"fyne.io/fyne/v2/widget"
)

func newStatusBar() *widget.Label {
	return widget.NewLabel("Open a file")
}
