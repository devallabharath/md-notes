package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// MainView :: main view of the app
func MainView(winSize fyne.Size, path fyne.URI) (*container.Split, *container.AppTabs) {
	tree := NewFileTree(path)
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
	mainview.SetOffset(float64(tree.MinSize().Width / winSize.Width))

	return mainview, tabs
}
