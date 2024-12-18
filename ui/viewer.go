package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/devallabharath/md-notes/core"
)

// MainView :: main view of the app
func MainView(winWidth float64) *container.Split {
	tree := NewFileTree(storage.NewFileURI(core.Config.Path))
	tabs := NewTabs()

	tree.OnSelected = func(uid widget.TreeNodeID) {
		uri := tree.uriCache[uid]
		if uri.Extension() != ".md" {
			return
		}
		tab, err := NewTab(uri)
		if err != nil {
			return
		}
		tabs.Append(tab.tabitem)
		tabs.Select(tab.tabitem)
	}

	mainview := container.NewHSplit(tree, tabs)
	mainview.SetOffset(float64(tree.MinSize().Width / 1000))

	return mainview
}
