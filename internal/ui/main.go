package ui

import (
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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

func (e *editor) fileSaveAs() {
	open := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, e.win)
			return
		}
		if write == nil {
			return
		}

		e.SaveAs(write)
	}, e.win)

	open.SetFilter(storage.NewExtensionFileFilter([]string{".png"}))
	open.Show()
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
		recents,
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Reset ...", e.fileReset),
		fyne.NewMenuItem("Save", e.fileSave),
		fyne.NewMenuItem("Save As ...", e.fileSaveAs),
	)

	return fyne.NewMainMenu(file)
}

func (e *editor) loadRecentMenu() *fyne.Menu {
	var items []*fyne.MenuItem
	for _, item := range e.loadRecent() {
		u := item
		label := filepath.Base(item.String())

		items = append(items, fyne.NewMenuItem(label, func() {
			read, err := storage.OpenFileFromURI(u)
			if err != nil {
				fyne.LogError("Unable to open file \""+u.String()+"\"", err)
				return
			}
			e.LoadFile(read)
		}))
	}

	if e.recentMenu == nil {
		e.recentMenu = fyne.NewMenu("Recent Items", items...)
	} else {
		e.recentMenu.Items = items
	}

	return e.recentMenu
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
