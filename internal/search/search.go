/*
	A bunch of helper functions for the search algorithms
*/

package search

import "search-visualizer/internal/grid"

// Neighbours returns the neighbours for a position
func Neighbours(position grid.Point) []*grid.Cell {
	neighbours := make([]*grid.Cell, 0)
	// Look at the (potential) 8 neighbours for this position
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if n := grid.GetCell(grid.Point{X: position.X + i, Y: position.Y + j}); n != nil {
				neighbours = append(neighbours, n)
			}
		}
	}
	return neighbours
}
