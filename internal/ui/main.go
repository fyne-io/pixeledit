package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/fyne-io/pixeledit/internal/api"
)

type toolbar struct {
	edit api.Editor
}

func (t *toolbar) toolbarSave() {
	t.edit.Save()
}

func (t *toolbar) toolbarReset() {
	t.edit.Reload()
}

func buildToolbar(e api.Editor) fyne.CanvasObject {
	t := &toolbar{edit: e}

	return widget.NewToolbar(
		&widget.ToolbarAction{Icon: theme.CancelIcon(), OnActivated: t.toolbarReset},
		&widget.ToolbarAction{Icon: theme.DocumentSaveIcon(), OnActivated: t.toolbarSave},
	)
}

// BuildUI creates the main window of our pixel edit application
func (e *editor) BuildUI() fyne.CanvasObject {
	palette := newPalette(e)
	toolbar := buildToolbar(e)

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, e.status, palette, nil),
		toolbar, e.status, palette, e.buildUI())
}
