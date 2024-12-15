package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	App := app.New()
	Window := App.NewWindow("MD-Notes")
	label := widget.NewLabel("Hello")
	Window.SetContent(label)
	Window.ShowAndRun()
}
