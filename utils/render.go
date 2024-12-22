package utils

import (
	"fyne.io/fyne/v2/widget"
)

// render :: render funcs
type render struct{}

// Cell :: returns cell
func (f *render) Cell(rune rune, style widget.TextGridStyle) widget.TextGridCell {
	cell := widget.TextGridCell{Rune: rune, Style: style}
	return cell
}

// EmptyCell :: returns cell
func (f *render) EmptyCell() widget.TextGridCell {
	cell := widget.TextGridCell{Rune: ' ', Style: widget.TextGridStyleDefault}
	return cell
}

// Row :: return row
func (f *render) Row(cells []widget.TextGridCell) widget.TextGridRow {
	row := widget.TextGridRow{Cells: cells, Style: widget.TextGridStyleDefault}
	return row
}

// Row :: return row
func (f *render) EmptyRow() widget.TextGridRow {
	row := widget.TextGridRow{Cells: []widget.TextGridCell{f.EmptyCell()}, Style: widget.TextGridStyleDefault}
	return row
}

// Text :: return header
func (f *render) Text(text string) []widget.TextGridCell {
	var cells []widget.TextGridCell
	for _, c := range text {
		cells = append(cells, f.Cell(c, MdStyles.Normal))
	}
	return cells
}

// Header :: return header
func (f *render) Header(text string) []widget.TextGridCell {
	var cells []widget.TextGridCell
	for _, c := range text {
		cells = append(cells, f.Cell(c, MdStyles.Heading))
	}
	return cells
}

// Render :: markdown renderer
var Render = render{}
