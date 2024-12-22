package utils

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

// GetFileContent :: gets the file content
func GetFileContent(uri fyne.URI) (string, error) {
	data, err := storage.Reader(uri)
	content, error := io.ReadAll(data)
	data.Close()
	if err != nil {
		return "", err
	}
	if error != nil {
		return "", error
	}

	return string(content), nil
}
