package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"github.com/devallabharath/md-notes/core"
)

// MainView :: main view of the app
func MainView() *fyne.Container {
	side := NewFileTree(storage.NewFileURI(core.Config.Path))
	return container.NewHBox(side)
}
