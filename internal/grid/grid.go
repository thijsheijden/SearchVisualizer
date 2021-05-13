package grid

// A Grid contains the cells and keeps track of the start and finish cells
type Grid struct {
	Cells            []*Cell  // The cells making up the grid
	start            Point    // The starting point
	finish           Point    // The finishing point
	CurrentlyPlacing CellType // The cell type we are currently placing down
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
	gridInstance = Grid{
		Cells: make([]*Cell, Rows*Columns, Rows*Columns),
	}
	for row := 0; row < Rows; row++ {
		for col := 0; col < Columns; col++ {
			cell := Cell{
				Tag:      new(bool),
				CellType: Empty,
			}
			gridInstance.Cells[col+(Columns*row)] = &cell
		}
	}
}
