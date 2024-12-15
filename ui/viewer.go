package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/devallabharath/md-notes/core"
)

// MainView :: main view of the app
func MainView() *container.Split {
	side := NewFileTree(storage.NewFileURI(core.Config.Path))
	tabs := container.NewAppTabs()
	label := widget.NewLabel("Label")
	tab := container.NewTabItem("test.md", label)
	tabs.Append(tab)

	return container.NewHSplit(side, tabs)
}
