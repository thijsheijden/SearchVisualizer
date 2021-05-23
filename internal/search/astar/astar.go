package astar

import (
	"container/heap"
	"log"
	"math"
	"search-visualizer/internal/grid"
	"search-visualizer/internal/search"
)

// Algorithm implements the A* algorithm
type Algorithm struct {
	priorityQueue priorityQueue
	shortestPaths map[*grid.Cell]*path
	visited       map[*grid.Cell]bool
	finish        grid.Point
}

type path struct {
	from *grid.Cell
	cost int
}

// Create creates a new instance of the A* algorithm
func Create() *Algorithm {
	// Create the Algorithm object, containing the priority queue
	current := Algorithm{
		priorityQueue: make(priorityQueue, 0),
		shortestPaths: make(map[*grid.Cell]*path),
		visited:       make(map[*grid.Cell]bool),
		finish:        grid.GetFinishCell(),
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
			if a.shortestPaths[prev].from.CellType != grid.Start {
				a.shortestPaths[prev].from.CellType = grid.Path
			}
			prev = a.shortestPaths[prev].from
			if prev == grid.GetStartCell() {
				break
			}
		}
		return true
	}

	// Go over all its neighbours
	for _, neighbour := range search.Neighbours(*closest.Cell.Position) {
		// Check if this cell can be visited
		if neighbour.CellType == grid.Wall {
			continue
		}

		// Calculate the new cost for this neighbour
		// Take the cost of the current node, add 1 for the move into the neighbour

		// Calculate the manhattan distance to the finish node
		manhattan := math.Abs(float64(neighbour.Position.X-a.finish.X)) + math.Abs(float64(neighbour.Position.Y-a.finish.Y))

		// Add the manhattan distance to the cost
		newCost := int(manhattan)

		log.Println(newCost)

		// Check if they have been visited
		if !a.visited[neighbour] {
			// No path to this cell yet, create one
			a.shortestPaths[neighbour] = &path{
				from: closest.Cell,
				cost: newCost,
			}
			heap.Push(&a.priorityQueue, &item{Cost: newCost, Cell: neighbour})
		}
	}

	a.visited[closest.Cell] = true
	if closest.Cell.CellType != grid.Start {
		closest.Cell.CellType = grid.Visited
	}

	// Check if there is still any unvisited cell
	if len(a.priorityQueue) == 0 {
		return true
	}
	return false
}

// Stop stops the algorithm
func (a *Algorithm) Stop() {

}
