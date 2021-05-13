package menu

import (
	"log"
	"search-visualizer/internal/grid"
	"strconv"

	"gioui.org/io/pointer"
	"gioui.org/widget"
)

// HandleInput handles input regarding the top menu
func HandleInput(gtx c) {
	handleGridSizeChange()
	handleCellPaintTypeChange(gtx)
}

func handleCellPaintTypeChange(gtx c) {
	for _, e := range gtx.Events(topMenu.cellPaintTag) {
		if _, ok := e.(pointer.Event); ok {
			// Change the cell type
			topMenu.cellPaintType++
			if topMenu.cellPaintType > 3 {
				topMenu.cellPaintType = 0
			}
			log.Println(topMenu.cellPaintType.String())
		}
	}
}

// Handle inputs for the number of columns and rows editors
func handleGridSizeChange() {
	for _, e := range topMenu.gridColumnsInput.Events() {
		switch e.(type) {
		case widget.ChangeEvent:
			if valid, val := checkInputIsValid(topMenu.gridColumnsInput.Text()); !valid {
				topMenu.gridColumnsInput.Delete(-1)
			} else {
				// Update the grid number of columns
				grid.Columns = val
				grid.New()
			}
		}
	}

	for _, e := range topMenu.gridRowsInput.Events() {
		switch e.(type) {
		case widget.ChangeEvent:
			if valid, val := checkInputIsValid(topMenu.gridRowsInput.Text()); !valid {
				topMenu.gridRowsInput.Delete(-1)
			} else {
				// Update the grid number of rows
				grid.Rows = val
				grid.New()
			}
		}
	}
}

func checkInputIsValid(input string) (bool, int) {
	// Check if the widget text is still a number, if not, delete the last character
	intValue, err := strconv.Atoi(input)
	if err != nil {
		// Delete the last character from the input field
		return false, 0
	}
	if intValue > 100 {
		return false, 0
	}
	return true, intValue
}
