package main

import (
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
	Window.SetContent(ui.MainView())
	Window.ShowAndRun()
}
