package graphics

import (
	"image/color"

	"gioui.org/layout"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

// Grid variables
var (
	gridRows         int     = 25
	gridColumns      int     = 25
	gridRowWeight    float32 = 1 / float32(gridRows)
	gridColumnWeight float32 = 1 / float32(gridColumns)
)

// Some NRGBA colors
var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
	cellColor  = color.NRGBA{R: 0xEC, G: 0xEC, B: 0xEC, A: 0xFF}
)
