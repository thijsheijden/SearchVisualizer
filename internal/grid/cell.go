package grid

import (
	"image/color"
	"log"

	"gioui.org/io/pointer"
)

// A Cell is a rectangle in the grid
type Cell struct {
	Tag      *bool    // The tag for this cell, used for pointer events
	CellType CellType // The type of this cell
}

// A Point is a simple x, y coordinate
type Point struct {
	x int
	y int
}

// CellType is used to know what type a cell is
type CellType int

const (
	// Empty denotes the empty cell
	Empty CellType = iota
	// Wall denotes a wall that the search algorithm can not go through
	Wall
	// Start denotes the starting point for the search algorithm
	Start
	// Finish denotes the goal for the search algo
	Finish
)

func (t CellType) String() string {
	return [...]string{"Empty", "Wall", "Start", "Finish"}[t]
}

// Possible cell colors
var (
	DefaultCellColor = color.NRGBA{R: 0xEC, G: 0xEC, B: 0xEC, A: 0xFF}
	FinishCellColor  = color.NRGBA{R: 0x77, G: 0xDD, B: 0x77, A: 0xFF}
	StartCellColor   = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	BlueCellColor    = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

// Reset resets all attributes of a cell
func (c *Cell) Reset() {
	c.CellType = Empty
}

// Clicked is triggered if a cell is simply clicked
func (c *Cell) Clicked(button pointer.Buttons) {
	switch button {
	case pointer.ButtonPrimary:
		log.Println("Left click")
	case pointer.ButtonSecondary:
		log.Println("Right click")
	}
}
