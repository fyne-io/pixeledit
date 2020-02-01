package main

import (
	"os"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"

	"github.com/fyne-io/pixeledit/internal/api"
	"github.com/fyne-io/pixeledit/internal/ui"
)

func loadFileArgs(e api.Editor) {
	if len(os.Args) < 2 {
		return
	}

	time.Sleep(time.Second / 3) // wait for us to be shown before loading so scales are correct
	e.LoadFile(os.Args[1])
}

func main() {
	e := ui.NewEditor()

	a := app.New()
	w := a.NewWindow("Pixel Editor")
	w.SetContent(e.BuildUI())
	w.Resize(fyne.NewSize(520, 320))

	go loadFileArgs(e)
	w.ShowAndRun()
}
