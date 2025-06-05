package themes

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type colors map[fyne.ThemeColorName]color.Color

type variants map[fyne.ThemeVariant]colors

var darkColors = colors{
	"transparent":     color.Transparent,
	"gray":            DarkPallette.bg2,
	"red":             DarkPallette.red,
	"green":           DarkPallette.green,
	"blue":            DarkPallette.blue,
	"yellow":          DarkPallette.yellow,
	"orange":          DarkPallette.orange,
	"cyan":            DarkPallette.cyan,
	"teal":            DarkPallette.teal,
	"lavender":        DarkPallette.lavender,
	"primary":         DarkPallette.blue,
	"focus":           DarkPallette.bg4,
	"hover":           DarkPallette.bg4,
	"selection":       DarkPallette.bg4,
	"border":          DarkPallette.bg3,
	"foreground":      DarkPallette.fg1,
	"text":            DarkPallette.fg1,
	"background":      DarkPallette.bg1,
	"scrollBar":       DarkPallette.fg2,
	"separator":       DarkPallette.fg2,
	"inputBackground": DarkPallette.bg2,
	"shadow":          DarkPallette.shadow,
}

var lightColors = colors{
	"transparent":       color.Transparent,
	"gray":              LightPallette.bg2,
	"red":               LightPallette.red,
	"green":             LightPallette.green,
	"blue":              LightPallette.blue,
	"yellow":            LightPallette.yellow,
	"orange":            LightPallette.orange,
	"cyan":              LightPallette.cyan,
	"teal":              LightPallette.teal,
	"lavender":          LightPallette.lavender,
	"text":              LightPallette.fg1,
	"foreground":        LightPallette.fg1,
	"background":        LightPallette.bg1,
	"focus":             LightPallette.bg4,
	"hover":             LightPallette.bg4,
	"primary":           LightPallette.blue,
	"selection":         LightPallette.bg4,
	"border":            LightPallette.bg3,
	"scrollBar":         LightPallette.fg2,
	"separator":         LightPallette.fg2,
	"overlayBackground": LightPallette.red,
	"inputBackground":   LightPallette.bg2,
	"inputBorder":       LightPallette.bg2,
	"shadow":            LightPallette.shadow,
}

// Colors :: general colors
var Colors = variants{0: darkColors, 1: lightColors}
