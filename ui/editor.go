package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/devallabharath/md-notes/utils"
)

// Editor :: type of the Editor
type Editor struct {
	writer  *widget.Entry
	preview *widget.TextGrid
}

// NewEditor :: creates new editor
func NewEditor(uri fyne.URI) (*Editor, error) {
	fileContent, err := utils.GetFileContent(uri)
	if err != nil {
		return nil, err
	}

	editor := widget.NewMultiLineEntry()
	editor.Wrapping = fyne.TextWrap(3)
	editor.TextStyle.TabWidth = 2
	editor.TextStyle.Monospace = true
	editor.SetText(fileContent)

	viewer := widget.NewTextGrid()
	viewer.SetText(fileContent)
	editor.OnChanged = viewer.SetText

	return &Editor{
		editor,
		viewer,
	}, nil
}
