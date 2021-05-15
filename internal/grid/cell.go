package grid

import (
	"image/color"

	"gioui.org/io/pointer"
)

// A Cell is a rectangle in the grid
type Cell struct {
	Tag      *bool    // The tag for this cell, used for pointer events
	CellType CellType // The type of this cell
	Position *Point   // The position of this cell in the grid
}

// A Point is a simple x, y coordinate
type Point struct {
	X int
	Y int
}

// CellType is used to know what type a cell is
type CellType int

const (
	// MARK: Paintable cell types

	// Wall denotes a wall that the search algorithm can not go through
	Wall CellType = iota
	// Start denotes the starting point for the search algorithm
	Start
	// Finish denotes the goal for the search algo
	Finish
	// Empty denotes the empty cell
	Empty

	// MARK: Non-paintable cell types

	// Visited denotes a cell that has been visited by the search algorithm
	Visited

	// Path denotes a cell that is part of the (current) shortest path
	Path
)

func (t CellType) String() string {
	return [...]string{"Wall", "Start", "Finish", "Empty"}[t]
}

// Possible cell colors
var (
	DefaultCellColor = color.NRGBA{R: 0xED, G: 0xED, B: 0xED, A: 0xFF}
	FinishCellColor  = color.NRGBA{R: 0x77, G: 0xDD, B: 0x77, A: 0xFF}
	StartCellColor   = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	BlueCellColor    = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}

	ShortestPathColor = color.NRGBA{R: 0x72, G: 0xd3, B: 0xfe, A: 0xFF}
	VisitedColor      = color.NRGBA{R: 0xD3, G: 0xD3, B: 0xD3, A: 0xFF}
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
			gridInstance.Cells[gridInstance.start.X+(Columns*gridInstance.start.Y)].CellType = Empty
		}

		// Set this cell as the start
		c.CellType = Start

		// Set the start point to the position of this cell
		gridInstance.start = c.Position
	case Finish:
		// Check if a finish point has already been set
		if gridInstance.finish != nil {
			// Remove the previous finish point
			gridInstance.Cells[gridInstance.finish.X+(Columns*gridInstance.finish.Y)].CellType = Empty
		}

		// Set this cell as the finish
		c.CellType = Finish

		// Set the finish point to the position of this cell
		gridInstance.finish = c.Position
	case Empty:
		c.reset()
	}
}
