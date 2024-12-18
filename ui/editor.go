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
	preview *widget.RichText
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
	editor.SetText(string(content))
	renderer := widget.NewRichTextFromMarkdown(string(content))
	editor.OnChanged = renderer.ParseMarkdown
	return &Editor{
		editor,
		renderer,
	}, nil
}
