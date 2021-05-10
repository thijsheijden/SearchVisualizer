package graphics

import (
	"strconv"

	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type inputFields struct {
	theme         *material.Theme
	nColumnsInput *widget.Editor
	nRowsInput    *widget.Editor
	testLabel     *widget.Label
}

func CreateGridSizeInputFields() {
	gridInputFields = inputFields{
		nColumnsInput: new(widget.Editor),
		nRowsInput:    new(widget.Editor),
		theme:         material.NewTheme(gofont.Collection()),
		testLabel:     new(widget.Label),
	}
}

var gridInputFields inputFields

func GridRowColumnInput(gtx C) D {
	return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(1, func(gtx C) D {
				return gridInputFields.displayInputField(gtx, gridInputFields.nRowsInput)
			}),
			layout.Flexed(1, func(gtx C) D {
				return gridInputFields.displayInputField(gtx, gridInputFields.nColumnsInput)
			}),
		)
	})
}

func (i inputFields) displayInputField(gtx C, input *widget.Editor) D {
	e := material.Editor(i.theme, input, "25")
	e.Editor.SingleLine = true
	border := widget.Border{
		Color:        red,
		CornerRadius: unit.Dp(4),
		Width:        unit.Dp(1),
	}
	return border.Layout(gtx, func(gtx C) D {
		return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
			return e.Layout(gtx)
		})
	})
}

// Handle inputs for the number of columns and rows editors
func HandleGridSizeChange() {
	for _, e := range gridInputFields.nColumnsInput.Events() {
		switch e.(type) {
		case widget.ChangeEvent:
			if valid, val := checkInputIsValid(gridInputFields.nColumnsInput.Text()); !valid {
				gridInputFields.nColumnsInput.Delete(-1)
			} else {
				// Update the grid number of columns
				gridColumns = val
				CreateNewGrid()
			}
		}
	}

	for _, e := range gridInputFields.nRowsInput.Events() {
		switch e.(type) {
		case widget.ChangeEvent:
			if valid, val := checkInputIsValid(gridInputFields.nRowsInput.Text()); !valid {
				gridInputFields.nRowsInput.Delete(-1)
			} else {
				// Update the grid number of rows
				gridRows = val
				CreateNewGrid()
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

func HandleCellClicks(q event.Queue) {
	for _, cell := range grid {
		for range q.Events(cell.tag) {
			cell.wall = !cell.wall
		}
	}
}
