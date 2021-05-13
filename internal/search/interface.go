package search

// A Algorithm should implement the following methods
type Algorithm interface {
	// Next generates the next step of the algorithm
	Next()

	// Stop stops the algorithm
	Stop()
}
