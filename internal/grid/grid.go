package grid

// A Grid contains the cells and keeps track of the start and finish cells
type Grid struct {
	Cells            []*Cell   // The cells making up the grid
	start            *Point    // The starting point
	finish           *Point    // The finishing point
	CurrentlyPlacing *CellType // The cell type we are currently placing down
	Dragging         bool      // Whether we are currently dragging
}

// Grid variables
var (
	Rows         int     = 25
	Columns      int     = 25
	RowWeight    float32 = 1 / float32(Rows)
	ColumnWeight float32 = 1 / float32(Columns)
)

// gridInstance is the used instance of a Grid
var gridInstance Grid

// New creates a new grid instance
func New() {
	// Check if there are already cells in the gridinstance
	if gridInstance.Cells != nil {
		for _, cell := range gridInstance.Cells {
			cell.Tag = nil
		}
	}

	gridInstance = Grid{
		Cells: make([]*Cell, Rows*Columns, Rows*Columns),
	}
	for row := 0; row < Rows; row++ {
		for col := 0; col < Columns; col++ {
			cell := Cell{
				Tag:      new(bool),
				CellType: Empty,
				Position: &Point{X: col, Y: row},
			}
			gridInstance.Cells[col+(Columns*row)] = &cell
		}
	}
}

// SetCellPaintType sets the type of cell that currently can be placed with the pointer
func SetCellPaintType(t *CellType) {
	gridInstance.CurrentlyPlacing = t
}

// GetCell gets a cell from the grid
func GetCell(x int, y int) *Cell {
	if x < Columns && y < Rows && x >= 0 && y >= 0 {
		return gridInstance.Cells[x+(Columns*y)]
	}
	return nil
}

// GetStartCell gets the starting cell
func GetStartCell() *Cell {
	if gridInstance.start != nil {
		return gridInstance.Cells[gridInstance.start.X+(Columns*gridInstance.start.Y)]
	}
	return gridInstance.Cells[1+(Columns*1)]
}

func GetFinishCell() Point {
	return *gridInstance.Cells[gridInstance.finish.X+(Columns*gridInstance.finish.Y)].Position
}

// Reset sets all cells back to default (except start, finish and walls)
func Reset() {
	for _, c := range gridInstance.Cells {
		if c.CellType == Visited || c.CellType == Path {
			c.CellType = Empty
		}
	}
}
