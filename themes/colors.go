package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type colors map[fyne.ThemeColorName]color.Color

type variants map[fyne.ThemeVariant]colors

var darkColors = colors{
	"white":       DarkPallette.white,
	"black":       DarkPallette.black,
	"transparent": color.Transparent,
	"red":         DarkPallette.red,
	"green":       DarkPallette.green,
	"blue":        DarkPallette.blue,
	"yellow":      DarkPallette.yellow,
	"orange":      DarkPallette.orange,
	"cyan":        DarkPallette.cyan,
	"teal":        DarkPallette.teal,
	"lavender":    DarkPallette.lavender,
	"primary":     DarkPallette.lavender,
	"focus":       DarkPallette.lavender,
	"foreground":  DarkPallette.white,
	"background":  color.NRGBA{0x11, 0x11, 0x1b, 0xff},
	"scrollBar":   color.NRGBA{0xef, 0xf1, 0xf5, 0xff},
	"shadow":      color.NRGBA{0xe6, 0xe9, 0xef, 0xff},
}

var lightColors = colors{
  "white":       LightPallette.white,
	"black":       LightPallette.black,
	"transparent": color.Transparent,
	"red":         LightPallette.red,
	"green":       LightPallette.green,
	"blue":        LightPallette.blue,
	"yellow":      LightPallette.yellow,
	"orange":      LightPallette.orange,
	"cyan":        LightPallette.cyan,
	"teal":        LightPallette.teal,
	"lavender":    LightPallette.lavender,
	"primary":     LightPallette.lavender,
	"foreground":  LightPallette.white,
	"background":  color.NRGBA{0xdc, 0xe0, 0xe8, 0xff},
	"scrollBar":   color.NRGBA{0x1e, 0x1e, 0x29, 0xff},
	"shadow":      color.NRGBA{0xe6, 0xe9, 0xef, 0xff},
}

// Colors :: general colors
var Colors = variants{1: lightColors, 2: darkColors}
