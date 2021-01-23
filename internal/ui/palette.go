package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/fyne-io/pixeledit/internal/api"
	"github.com/fyne-io/pixeledit/internal/tool"
)

type palette struct {
	edit *editor

	zoom *widget.Label
}

func (p *palette) updateZoom(val int) {
	if val < 1 {
		val = 1
	} else if val > 16 {
		val = 16
	}
	p.edit.setZoom(val)

	p.zoom.SetText(fmt.Sprintf("%d%%", p.edit.zoom*100))
}

func (p *palette) loadTools() []api.Tool {
	return []api.Tool{
		&tool.Picker{},
		&tool.Pencil{},
	}
}

func newPalette(edit *editor) fyne.CanvasObject {
	p := &palette{edit: edit, zoom: widget.NewLabel("100%")}

	tools := p.loadTools()

	var toolIcons []fyne.CanvasObject
	for _, item := range tools {
		var icon *widget.Button
		thisTool := item
		icon = widget.NewButtonWithIcon(item.Name(), item.Icon(), func() {
			for _, toolButton := range toolIcons {
				toolButton.(*widget.Button).Importance = widget.MediumImportance
				toolButton.Refresh()
			}
			icon.Importance = widget.HighImportance
			icon.Refresh()
			edit.setTool(thisTool)
		})
		toolIcons = append(toolIcons, icon)
	}

	edit.setTool(tools[0])
	toolIcons[0].(*widget.Button).Importance = widget.HighImportance

	zoom := container.NewHBox(
		widget.NewButtonWithIcon("", theme.ZoomOutIcon(), func() {
			p.updateZoom(p.edit.zoom / 2)
		}),
		p.zoom,
		widget.NewButtonWithIcon("", theme.ZoomInIcon(), func() {
			p.updateZoom(p.edit.zoom * 2)
		}))

	return container.NewVBox(append([]fyne.CanvasObject{container.NewGridWithColumns(1),
		widget.NewLabel("Tools"), zoom, edit.fgPreview}, toolIcons...)...)
}
