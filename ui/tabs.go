package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// Tab :: tab details and func(save, saveas)
type Tab struct {
	tabitem *container.TabItem
	Editor  *Editor
	uri     fyne.URI
}

func (tab *Tab) save() {
	fmt.Println(tab.uri.Name(), "saved")
	content := tab.Editor.writer.Text
	fmt.Println(content)
}

// NewTabs :: Creates new tabs component
func NewTabs() *container.AppTabs {
	return container.NewAppTabs()
}

// NewTab :: Creates new tab
func NewTab(uri fyne.URI) (*Tab, error) {
	Editor, err := NewEditor(uri)
	if err != nil {
		return nil, err
	}
	editSroll := container.NewScroll(Editor.writer)
	previewSroll := container.NewScroll(Editor.preview)
	view := container.NewHSplit(editSroll, previewSroll)
	tab := container.NewTabItem(uri.Name(), view)

	return &Tab{
		tab,
		Editor,
		uri,
	}, nil
}
