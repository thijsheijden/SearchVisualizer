package grid

import (
	"image/color"

	"gioui.org/io/pointer"
)

// A Cell is a rectangle in the grid
type Cell struct {
	Tag    *bool // The tag for this cell, used for pointer events
	Wall   bool  // Whether this cell is currently a wall
	Start  bool  // Whether this cell is currently the starting point
	Finish bool  // Whether this cell is currently the finish point
}

// A Point is a simple x, y coordinate
type Point struct {
	x int
	y int
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
	c.Finish = false
	c.Start = false
	c.Wall = false
}

func (c *Cell) Clicked(button pointer.Buttons) {
	switch button {
	// Make the cell a wall
	case pointer.ButtonPrimary:
		prevWall := c.Wall
		c.Reset()
		c.Wall = !prevWall
	// Make the cell the finish or start
	case pointer.ButtonSecondary:
		// If this cell is currently the start, make it the finish
		if c.Start {
			c.Start = false
			c.Finish = true
		} else if c.Finish {
			// Reset the cell
			c.Reset()
		} else {
			// Make this cell the start
			c.Reset()
			c.Start = true
		}
	}
}
