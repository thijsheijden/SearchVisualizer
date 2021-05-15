/*
	A bunch of helper functions for the search algorithms
*/

package search

import "search-visualizer/internal/grid"

// Neighbours returns the neighbours for a position
func Neighbours(position grid.Point) []*grid.Cell {
	neighbours := make([]*grid.Cell, 0)
	// Get the 4 possible neighbours (non-diagonals)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Make sure this is not the cell we are, or a diagonal cell
			if !(i == 0 && j == 0) && !(i != 0 && j != 0) {
				if n := grid.GetCell(position.X+i, position.Y+j); n != nil {
					neighbours = append(neighbours, n)
				}
			}
		}
	}
	return neighbours
}
