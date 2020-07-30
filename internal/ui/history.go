package ui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/storage"
)

const (
	recentCountKey  = "recentCount"
	recentFormatKey = "recent-%d"
)

func (e *editor) loadRecent() []fyne.URI {
	pref := fyne.CurrentApp().Preferences()
	count := pref.Int(recentCountKey)

	recent := []fyne.URI{}
	for i := 0; i < count; i++ {
		key := fmt.Sprint(recentFormatKey, i)
		recent = append(recent, storage.NewURI(pref.String(key)))
	}

	return recent
}

func (e *editor) addRecent(u fyne.URI) {
	pref := fyne.CurrentApp().Preferences()
	recent := e.loadRecent()

	recent = append([]fyne.URI{u}, recent...)
	if len(recent) > 5 {
		recent = recent[:5]
	}

	pref.SetInt(recentCountKey, len(recent))
	for i, uri := range recent {
		key := fmt.Sprint(recentFormatKey, i)
		pref.SetString(key, uri.String())
	}

	e.loadRecentMenu()
	e.win.SetMainMenu(e.win.MainMenu())
}
