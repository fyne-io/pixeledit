package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/fyne-io/pixeledit/internal/api"
)

type toolbar struct {
	edit api.Editor
}

func (t *toolbar) toolbarOpen() {
	win := fyne.CurrentApp().Driver().AllWindows()[0]
	open := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if read == nil {
			return
		}

		t.edit.LoadFile(read)
	}, win)

	open.SetFilter(storage.NewExtensionFileFilter([]string{".png"}))
	open.Show()
}

func (t *toolbar) toolbarReset() {
	win := fyne.CurrentApp().Driver().AllWindows()[0]
	dialog.ShowConfirm("Reset content?", "Are you sure you want to re-load the image?",
		func(ok bool) {
			if !ok {
				return
			}

			t.edit.Reload()
		}, win)
}

func (t *toolbar) toolbarSave() {
	t.edit.Save()
}

func buildToolbar(e api.Editor) fyne.CanvasObject {
	t := &toolbar{edit: e}

	return widget.NewToolbar(
		&widget.ToolbarAction{Icon: theme.FolderOpenIcon(), OnActivated: t.toolbarOpen},
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
