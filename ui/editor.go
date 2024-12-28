package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/devallabharath/md-notes/markdown"
	"github.com/devallabharath/md-notes/utils"
)

// Editor :: type of the Editor
type Editor struct {
	editor  *widget.Entry
	viewer *markdown.Md
}

// NewEditor :: creates new editor
func NewEditor(uri fyne.URI) (*Editor, error) {
	var content []byte
	var err error

	if content, err = utils.GetFileContent(uri); err != nil {
		return nil, err
	}

	editor := widget.NewMultiLineEntry()
	editor.Wrapping = fyne.TextWrap(3)
	editor.TextStyle.TabWidth = 2
	editor.TextStyle.Monospace = true
	editor.SetText(string(content))

	viewer := markdown.New(uri, content)
	viewer.Wrapping = fyne.TextWrap(3)

	// editor.OnChanged = func(content string) {}

	return &Editor{
		editor,
		viewer,
	}, nil
}
