package ui

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// Editor :: type of the Editor
type Editor struct {
	writer  *widget.Entry
	preview *widget.TextGrid
}

// NewEditor :: creates new editor
func NewEditor(uri fyne.URI) (*Editor, error) {
	data, err := storage.Reader(uri)
	content, error := io.ReadAll(data)
	data.Close()
	if err != nil {
		return nil, err
	}
	if error != nil {
		return nil, error
	}
	editor := widget.NewMultiLineEntry()
	editor.Wrapping = fyne.TextWrap(3)
	editor.TextStyle.TabWidth = 2
	editor.TextStyle.Monospace = true
	editor.SetText(string(content))
	viewer := widget.NewTextGrid()
	viewer.SetText(string(content))
	editor.OnChanged = viewer.SetText
	return &Editor{
		editor,
		viewer,
	}, nil
}
