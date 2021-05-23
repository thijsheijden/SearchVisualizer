package menu

import (
	"image/color"
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

	// Whether the algorithm is running
	algorithmRunning bool

	// The type of cell we are currently painting
	cellPaintTag       *bool
	cellPaintType      grid.CellType
	cellPaintTypeLabel *widget.Label

	// Start and reset buttons
	startPauseButton *widget.Clickable
	resetButton      *widget.Clickable

	// Algorithm control func in main
	algoControlFunc func(command ButtonCommand)
}

var topMenu menu

// New creates a new top menu
func New(algoControlFunc func(command ButtonCommand)) {
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
		startPauseButton:      new(widget.Clickable),
		resetButton:           new(widget.Clickable),
		algoControlFunc:       algoControlFunc,
	}
}

var (
	startButtonColor = color.NRGBA{R: 0x77, G: 0xDD, B: 0x77, A: 0xFF}
	pauseButtonColor = color.NRGBA{R: 0xFF, G: 0xB3, B: 0x47, A: 0xFF}
	resetButtonColor = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
)

// PassCellTypeToGrid passes the memory address of the celltype in the menu to the grid
func PassCellTypeToGrid() {
	grid.SetCellPaintType(&topMenu.cellPaintType)
}

// ToggleAlgorithmRunning toggles whether the menu thinks the algorithm is running
func ToggleAlgorithmRunning() {
	topMenu.algorithmRunning = !topMenu.algorithmRunning
}
