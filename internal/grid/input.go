package grid

import "gioui.org/io/pointer"

// HandleInput handles input on the grid (i.e. cell clicks)
func HandleInput(gtx c) {
	for _, cell := range gridInstance.Cells {
		for _, e := range gtx.Events(cell.Tag) {
			// Check if it is a pointer event
			if e, ok := e.(pointer.Event); ok {
				cell.Clicked(e.Buttons)
			}
		}
	}
}
