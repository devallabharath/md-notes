package markdown

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var (
	// text styles
	Normal          = fyne.TextStyle{}
	Bold            = fyne.TextStyle{Bold: true}
	Italic          = fyne.TextStyle{Italic: true}
	BoldItalic      = fyne.TextStyle{Bold: true, Italic: true}
	Underline       = fyne.TextStyle{Underline: true}
	BoldUnderline   = fyne.TextStyle{Bold: true, Underline: true}
	ItalicUnderline = fyne.TextStyle{Italic: true, Underline: true}
	// text alignment
	Left   = fyne.TextAlign(0)
	Center = fyne.TextAlign(1)
	Right  = fyne.TextAlign(2)
	// text size
	Bigger  = fyne.ThemeSizeName("headingText")
	Big     = fyne.ThemeSizeName("subHeadingText")
	Small   = fyne.ThemeSizeName("text")
	Smaller = fyne.ThemeSizeName("smaller")
	Tiny    = fyne.ThemeSizeName("tiny")
	// text colors
	Text   = fyne.ThemeColorName("text")
	Red    = fyne.ThemeColorName("red")
	Green  = fyne.ThemeColorName("green")
	Blue   = fyne.ThemeColorName("blue")
	Yellow = fyne.ThemeColorName("yellow")
	Orange = fyne.ThemeColorName("orange")
)

type mdstyles struct {
	Text       widget.RichTextStyle
	Paragraph  widget.RichTextStyle
	Heading    widget.RichTextStyle
	Heading1   widget.RichTextStyle
	Heading2   widget.RichTextStyle
	Emphasis   widget.RichTextStyle
	Italic     widget.RichTextStyle
	Blockquote widget.RichTextStyle
	Codeblock  widget.RichTextStyle
	Emptyline  widget.RichTextStyle
}

// MdStyles :: markdown styles
var MdStyles = &mdstyles{
	Text: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Text,
		Inline:    false,
		SizeName:  Small,
		TextStyle: Normal,
	},
	Paragraph: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Text,
		Inline:    false,
		SizeName:  Small,
		TextStyle: Normal,
	},
	Heading: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Text,
		Inline:    false,
		SizeName:  Big,
		TextStyle: Bold,
	},
	Heading1: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Yellow,
		Inline:    false,
		SizeName:  Bigger,
		TextStyle: Bold,
	},
	Heading2: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Yellow,
		Inline:    false,
		SizeName:  Bigger,
		TextStyle: Bold,
	},
	Emphasis: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Text,
		Inline:    false,
		SizeName:  Small,
		TextStyle: Bold,
	},
	Italic: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Text,
		Inline:    false,
		SizeName:  Small,
		TextStyle: Italic,
	},
	Blockquote: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Text,
		Inline:    false,
		SizeName:  Small,
		TextStyle: Normal,
	},
	Codeblock: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Text,
		Inline:    false,
		SizeName:  Small,
		TextStyle: Normal,
	},
	Emptyline: widget.RichTextStyle{
		Alignment: Left,
		ColorName: Text,
		Inline:    false,
		SizeName:  Tiny,
		TextStyle: Normal,
	},
}
