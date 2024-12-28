package markdown

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/devallabharath/md-notes/utils"
)

var colors = utils.Colors

type styles struct {
	Normal    fyne.TextStyle
	Bold      fyne.TextStyle
	Italic    fyne.TextStyle
	Underline fyne.TextStyle
}

type gridStyle struct {
	style fyne.TextStyle
	fg    color.Color
	bg    color.Color
}

func (c *gridStyle) Style() fyne.TextStyle {
	return c.style
}
func (c *gridStyle) TextColor() color.Color {
	return c.fg
}
func (c *gridStyle) BackgroundColor() color.Color {
	return c.bg
}

type mdstyles struct {
	Normal     widget.TextGridStyle
	Heading    widget.TextGridStyle
	Emphasis   widget.TextGridStyle
	Italic     widget.TextGridStyle
	Blockquote widget.TextGridStyle
}

// TextStyles :: text styles like bold, italic, underline
var TextStyles = &styles{
	fyne.TextStyle{},
	fyne.TextStyle{Bold: true},
	fyne.TextStyle{Italic: true},
	fyne.TextStyle{Underline: true},
}

// MdStyles :: markdown styles
var MdStyles = &mdstyles{
	&gridStyle{TextStyles.Bold, colors.White, colors.Transparent},
	&gridStyle{TextStyles.Bold, colors.Yellow, colors.Transparent},
	&gridStyle{TextStyles.Bold, colors.White, colors.Transparent},
	&gridStyle{TextStyles.Bold, colors.White, colors.Transparent},
	&gridStyle{TextStyles.Bold, colors.White, colors.Black},
}
