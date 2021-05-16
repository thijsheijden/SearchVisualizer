package grid

import (
	"log"

	"gioui.org/io/pointer"
)

// HandleInput handles input on the grid (i.e. cell clicks)
func HandleInput(gtx c) {
	for _, cell := range gridInstance.Cells {
		for _, e := range gtx.Events(cell.Tag) {
			// Check if it is a pointer event
			if e, ok := e.(pointer.Event); ok {
				switch e.Type {
				case pointer.Press:
					gridInstance.Dragging = true
					cell.Clicked(e.Buttons)
				case pointer.Release:
					gridInstance.Dragging = false
				case pointer.Enter:
					log.Println("Entered")
					log.Println(cell.Tag)
					// If we are currently dragging, color this cell
					if gridInstance.Dragging {
						cell.set()
					}
				}

			}
		}
	}
}
