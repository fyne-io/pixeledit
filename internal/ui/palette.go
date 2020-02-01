package ui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

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
				toolButton.(*widget.Button).Style = widget.DefaultButton
				toolButton.Refresh()
			}
			icon.Style = widget.PrimaryButton
			icon.Refresh()
			edit.setTool(thisTool)
		})
		toolIcons = append(toolIcons, icon)
	}

	edit.setTool(tools[0])
	toolIcons[0].(*widget.Button).Style = widget.PrimaryButton

	zoom := widget.NewHBox(
		widget.NewButtonWithIcon("", theme.ZoomOutIcon(), func() {
			p.updateZoom(p.edit.zoom / 2)
		}),
		p.zoom,
		widget.NewButtonWithIcon("", theme.ZoomInIcon(), func() {
			p.updateZoom(p.edit.zoom * 2)
		}))

	return widget.NewVBox(append([]fyne.CanvasObject{fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		widget.NewLabel("Tools"), zoom, edit.fgPreview)}, toolIcons...)...)
}
