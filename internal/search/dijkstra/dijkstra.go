package dijkstra

import (
	"container/heap"
	"log"
	"search-visualizer/internal/grid"
	"search-visualizer/internal/search"
)

// Algorithm implements Dijkstra's shortest path algorithm
type Algorithm struct {
	priorityQueue priorityQueue
	shortestPaths map[*grid.Cell]*path
	visited       map[*grid.Cell]bool
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
		shortestPaths: make(map[*grid.Cell]*path),
		visited:       make(map[*grid.Cell]bool),
	}

	// Push the starting node into the heap with cost 0
	heap.Push(&current.priorityQueue, &item{Cost: 0, Cell: grid.GetStartCell()})

	return &current
}

// Next performs the next algorithm step
// Returns a boolean denoting whether the algo is finished
func (a *Algorithm) Next() bool {
	// Get the closest node from the heap
	closest := heap.Pop(&a.priorityQueue).(*item)

	// If this is the finish node, stop
	if closest.Cell.CellType == grid.Finish {
		log.Println("Found finish")
		var prev *grid.Cell
		prev = closest.Cell
		for {
			a.shortestPaths[prev].from.CellType = grid.Path
			prev = a.shortestPaths[prev].from
			if prev == grid.GetStartCell() {
				break
			}
		}
		return true
	}

	newCost := closest.Cost + 1

	// Go over all its neighbours
	for _, neighbour := range search.Neighbours(*closest.Cell.Position) {
		// Check if this cell can be visited
		if neighbour.CellType == grid.Wall {
			continue
		}

		// Check if they have been visited
		if !a.visited[neighbour] {
			// If not, check if this new path is shorter than the one in the queue (if there is already a path)
			if c, ok := a.shortestPaths[neighbour]; ok {
				// Check if this path is shorter
				if newCost < c.cost {
					// Update path
					c.cost = newCost
					c.from = closest.Cell
				}

				// TODO: Update this in the heap
			} else {
				// No path to this cell yet, create one
				a.shortestPaths[neighbour] = &path{
					from: closest.Cell,
					cost: newCost,
				}
				heap.Push(&a.priorityQueue, &item{Cost: newCost, Cell: neighbour})
			}
		}
	}

	a.visited[closest.Cell] = true
	closest.Cell.CellType = grid.Visited
	return false
}

// Stop stops the algorithm
func (a *Algorithm) Stop() {

}
