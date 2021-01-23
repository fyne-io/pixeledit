package main

import (
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/storage"

	"github.com/fyne-io/pixeledit/internal/api"
	"github.com/fyne-io/pixeledit/internal/ui"
)

func loadFileArgs(e api.Editor) {
	if len(os.Args) < 2 {
		return
	}

	time.Sleep(time.Second / 3) // wait for us to be shown before loading so scales are correct
	read, err := storage.Reader(storage.NewFileURI(os.Args[1]))
	if err != nil {
		fyne.LogError("Unable to open file \""+os.Args[1]+"\"", err)
		return
	}
	e.LoadFile(read)
}

func main() {
	e := ui.NewEditor()

	a := app.NewWithID("io.fyne.pixeledit")
	w := a.NewWindow("Pixel Editor")
	e.BuildUI(w)
	w.Resize(fyne.NewSize(520, 320))

	go loadFileArgs(e)
	w.ShowAndRun()
}
