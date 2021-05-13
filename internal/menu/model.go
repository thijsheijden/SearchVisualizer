package menu

import (
	"search-visualizer/internal/grid"

	"gioui.org/font/gofont"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type menu struct {
	theme                 *material.Theme
	gridColumnsInput      *widget.Editor
	gridColumnsInputLabel *widget.Label
	gridRowsInput         *widget.Editor
	gridRowsInputLabel    *widget.Label
	tickSpeedInput        *widget.Editor
	tickSpeedInputLabel   *widget.Label

	// The type of cell we are currently painting
	cellPaintTag       *bool
	cellPaintType      grid.CellType
	cellPaintTypeLabel *widget.Label
}

var topMenu menu

// New creates a new top menu
func New() {
	topMenu = menu{
		theme:                 material.NewTheme(gofont.Collection()),
		gridColumnsInput:      new(widget.Editor),
		gridColumnsInputLabel: new(widget.Label),
		gridRowsInput:         new(widget.Editor),
		gridRowsInputLabel:    new(widget.Label),
		tickSpeedInput:        new(widget.Editor),
		tickSpeedInputLabel:   new(widget.Label),
		cellPaintTag:          new(bool),
		cellPaintType:         grid.Wall,
		cellPaintTypeLabel:    new(widget.Label),
	}
}
