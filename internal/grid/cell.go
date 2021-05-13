package grid

import (
	"image/color"

	"gioui.org/io/pointer"
)

// A Cell is a rectangle in the grid
type Cell struct {
	Tag      *bool    // The tag for this cell, used for pointer events
	CellType CellType // The type of this cell
	position *Point   // The position of this cell in the grid
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

// reset resets all attributes of a cell
func (c *Cell) reset() {
	// Check if this is a start or finish cell
	if c.CellType == Start {
		gridInstance.start = nil
	} else if c.CellType == Finish {
		gridInstance.finish = nil
	}
	c.CellType = Empty
}

// Clicked is triggered if a cell is simply clicked
func (c *Cell) Clicked(button pointer.Buttons) {
	switch button {
	case pointer.ButtonPrimary:
		c.set()
	case pointer.ButtonSecondary:
		c.reset()
	}
}

// set sets the cell to the currently selected type
func (c *Cell) set() {
	switch *gridInstance.CurrentlyPlacing {
	case Wall:
		c.CellType = Wall
	case Start:
		// Check if a start point has already been set
		if gridInstance.start != nil {
			// Remove the previous start point
			gridInstance.Cells[gridInstance.start.x+(Columns*gridInstance.start.y)].CellType = Empty
		}

		// Set this cell as the start
		c.CellType = Start

		// Set the start point to the position of this cell
		gridInstance.start = c.position
	case Finish:
		// Check if a finish point has already been set
		if gridInstance.finish != nil {
			// Remove the previous finish point
			gridInstance.Cells[gridInstance.finish.x+(Columns*gridInstance.finish.y)].CellType = Empty
		}

		// Set this cell as the finish
		c.CellType = Finish

		// Set the finish point to the position of this cell
		gridInstance.finish = c.position
	case Empty:
		c.reset()
	}
}
