package utils

import "image/color"

type colors struct {
	White       color.Color
	Black       color.Color
	Transparent color.Color
	Red         color.Color
	Green       color.Color
	Blue        color.Color
	Yellow      color.Color
	Orange      color.Color
	Cyan        color.Color
}

// Colors :: general colors
var Colors = &colors{
	color.White,
	color.Black,
	color.Transparent,
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{255, 255, 0, 255},
	color.RGBA{255, 255, 0, 255},
	color.RGBA{0, 255, 255, 255},
}
