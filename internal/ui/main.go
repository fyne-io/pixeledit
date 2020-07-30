package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func (e *editor) fileOpen() {
	open := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, e.win)
			return
		}
		if read == nil {
			return
		}

		e.LoadFile(read)
	}, e.win)

	open.SetFilter(storage.NewExtensionFileFilter([]string{".png"}))
	open.Show()
}

func (e *editor) fileReset() {
	win := fyne.CurrentApp().Driver().AllWindows()[0]
	dialog.ShowConfirm("Reset content?", "Are you sure you want to re-load the image?",
		func(ok bool) {
			if !ok {
				return
			}

			e.Reload()
		}, win)
}

func (e *editor) fileSave() {
	e.Save()
}


func buildToolbar(e *editor) fyne.CanvasObject {
	return widget.NewToolbar(
		&widget.ToolbarAction{Icon: theme.FolderOpenIcon(), OnActivated: e.fileOpen},
		&widget.ToolbarAction{Icon: theme.CancelIcon(), OnActivated: e.fileReset},
		&widget.ToolbarAction{Icon: theme.DocumentSaveIcon(), OnActivated: e.fileSave},
	)
}

func (e *editor) buildMainMenu() *fyne.MainMenu {
	recents := fyne.NewMenuItem("Open Recent", nil)
	recents.ChildMenu = e.loadRecentMenu()

	file := fyne.NewMenu("File",
		fyne.NewMenuItem("Open ...", e.fileOpen),
		fyne.NewMenuItem("Reset ...", e.fileReset),
		fyne.NewMenuItem("Save", e.fileSave),
	)

	return fyne.NewMainMenu(file)
}
// BuildUI creates the main window of our pixel edit application
func (e *editor) BuildUI(w fyne.Window) {
	palette := newPalette(e)
	toolbar := buildToolbar(e)
	e.win = w
	w.SetMainMenu(e.buildMainMenu())

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, e.status, palette, nil),
		toolbar, e.status, palette, e.buildUI())
	w.SetContent(content)
}
