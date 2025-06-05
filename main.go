package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/devallabharath/md-notes/core"
	"github.com/devallabharath/md-notes/themes"
	"github.com/devallabharath/md-notes/ui"
)

type appType struct {
	App     fyne.App
	Window  fyne.Window
	Menu    *widget.Menu
	Toolbar *widget.Toolbar
	Sidebar *ui.FileTree
	AppTabs *container.AppTabs
	Tabs    []*container.TabItem
	View    *container.Split
}

// App :: runtime instance with all objects
var App *appType

func init() {
	core.NewConfig()
	App = &appType{}
	App.App = app.New()
	App.App.Settings().SetTheme(themes.NewTheme(fyne.ThemeVariant(2)))
}

func newTab(file fyne.URI) {
	tab, _ := ui.NewTab(file)
	App.AppTabs.Append(tab.Tab)
}

// NewWindow :: creates a new window and returna
func NewWindow(title string, size fyne.Size, path fyne.URI) {
	win := App.App.NewWindow(title)
	win.Resize(size)
	App.Window = win
	//App.View, App.Sidebar, App.AppTabs = ui.MainView(size, path)
	//App.Window.SetContent(App.View)
	App.AppTabs = ui.MainView(size, path)
	App.Window.SetContent(App.AppTabs)
	if len(os.Args) > 1 {
		url := os.Args[1]
		newTab(storage.NewFileURI(url))
	} else {
		newTab(storage.NewFileURI("./test.md"))
	}
	// App.Sidebar.Hide()
	win.ShowAndRun()
}

func main() {
	winSize := fyne.NewSize(1000, 600)
	dir, _ := os.UserHomeDir()
	path := storage.NewFileURI(dir)
	NewWindow("MD-Notes", winSize, path)
}
