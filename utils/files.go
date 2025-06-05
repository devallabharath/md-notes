package utils

import (
	"io"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

// GetFileContent :: gets the file content
func GetFileContent(uri fyne.URI) ([]byte, error) {
	var data fyne.URIReadCloser
	var bytes []byte
	var err error
	if data, err = storage.Reader(uri); err != nil {
		return nil, err
	}
	if bytes, err = io.ReadAll(data); err != nil {
		return nil, err
	}
	data.Close()

	return bytes, nil
}

// GetImageURI :: gets absolute path of the image
func GetImageURI(URI fyne.URI, rpath string) fyne.URI {
	dir := filepath.Dir(URI.Path())
	imgPath := filepath.Join(dir, rpath)
	return storage.NewFileURI(imgPath)
}
