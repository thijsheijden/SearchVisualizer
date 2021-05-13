package dijkstra

import (
	"container/heap"
	"search-visualizer/internal/grid"
	"search-visualizer/internal/search"
)

// Algorithm implements Dijkstra's shortest path algorithm
type Algorithm struct {
	priorityQueue priorityQueue
	shortestPaths map[*grid.Cell]path
}

type path struct {
	from *grid.Cell
	cost int
}

// Create creates a new instance of Dijkstra's algorithm
func Create() *Algorithm {
	// Create the Algorithm object, containing the priority queue
	current := Algorithm{
		priorityQueue: make(priorityQueue, 0),
		shortestPaths: make(map[*grid.Cell]path),
	}

	// Push the start node to the priority queue with a cost of 0
	heap.Push(&current.priorityQueue, &item{Cost: 0, Cell: grid.GetCell(grid.Point{X: 1, Y: 1})})

	return &current
}

// Next performs the next algorithm step
func (a *Algorithm) Next() {
	// Get the closest node from the priority queue
	top := heap.Pop(&a.priorityQueue).(*item)

	// Mark this node as visited
	if !(top.Cell.CellType == grid.Start) {
		top.Cell.CellType = grid.Visited
	}

	// Get the neighbours for this node
	for _, cell := range search.Neighbours(*top.Cell.Position) {
		// Check if we already have a distance for these cells
		if p, ok := a.shortestPaths[cell]; ok {
			// Check if the new path is shorter
			if p.cost < top.Cost+1 {
				// If this is the case, update the shortest path
				p := a.shortestPaths[cell]
				p.cost = top.Cost + 1
				a.shortestPaths[cell] = p
			}
		}
	}
}

// Stop stops the algorithm
func (a *Algorithm) Stop() {

}
