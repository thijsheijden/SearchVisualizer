package grid

// A Grid contains the cells and keeps track of the start and finish cells
type Grid struct {
	Cells  []*Cell // The cells making up the grid
	start  Point   // The starting point
	finish Point   // The finishing point
}

// Grid variables
var (
	Rows         int     = 25
	Columns      int     = 25
	RowWeight    float32 = 1 / float32(Rows)
	ColumnWeight float32 = 1 / float32(Columns)
)

// GridInstance is the used instance of a Grid
var GridInstance Grid

// New creates a new grid instance
func New() {
	GridInstance = Grid{
		Cells: make([]*Cell, Rows*Columns, Rows*Columns),
	}
	for r := 0; r < Rows; r++ {
		for c := 0; c < Columns; c++ {
			cell := Cell{
				Tag:    new(bool),
				Wall:   false,
				Start:  false,
				Finish: false,
			}
			GridInstance.Cells[c+(Columns*r)] = &cell
		}
	}
}
