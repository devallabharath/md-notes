package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/devallabharath/md-notes/core"
	"github.com/devallabharath/md-notes/ui"
)

func init() {
	core.NewConfig()
}

func main() {
	App := app.New()
	Window := App.NewWindow("MD-Notes")
	Window.Resize(fyne.NewSize(1000, 600))
	Window.SetContent(ui.MainView(1000))
	Window.ShowAndRun()
}
