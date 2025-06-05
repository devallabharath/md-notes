package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// MainView :: main view of the app
func MainView(winSize fyne.Size, path fyne.URI) *container.AppTabs {
	return NewTabs()
}

// MainView :: main view of the app
//func MainView(winSize fyne.Size, path fyne.URI) (*container.Split, *FileTree, *container.AppTabs) {
//	tree := NewFileTree(path)
//	tabs := NewTabs()
//
//	tree.OnSelected = func(uid widget.TreeNodeID) {
//		uri := tree.uriCache[uid]
//		if uri.Extension() != ".md" {
//			return
//		}
//		tab, err := NewTab(uri)
//		if err != nil {
//			return
//		}
//		tabs.Append(tab.Tab)
//		tabs.Select(tab.Tab)
//	}
//
//	mainview := container.NewHSplit(tree, tabs)
//	mainview.SetOffset(float64(tree.MinSize().Width / winSize.Width))
//
//	return mainview, tree, tabs
//}
