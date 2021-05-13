package grid

import (
	"image"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type (
	c = layout.Context
	d = layout.Dimensions
)

func Display(gtx c) d {
	var rows []layout.FlexChild
	for row := 0; row < Rows; row++ {
		f := drawRowFunc(gtx, row)
		rows = append(rows,
			layout.Flexed(RowWeight, func(gtx c) d {
				return f(gtx)
			}),
		)
	}
	return layout.Inset{
		Top:    unit.Dp(0),
		Left:   unit.Dp(8),
		Right:  unit.Dp(8),
		Bottom: unit.Dp(8),
	}.Layout(gtx, func(gtx c) d {
		return layout.Flex{Axis: layout.Vertical, Alignment: layout.Alignment(layout.Center)}.Layout(gtx, rows...)
	})
}

func drawRowFunc(gtx c, row int) func(c) d {
	return func(gtx c) d {
		return drawRow(gtx, row)
	}
}

func drawBoxFunc(gtx c, size image.Point, row int, col int) func(c) d {
	return func(gtx c) d {
		return colorBox(gtx, gtx.Constraints.Max, row, col)
	}
}

func drawRow(gtx c, row int) d {
	var columns []layout.FlexChild
	for col := 0; col < Columns; col++ {
		f := drawBoxFunc(gtx, gtx.Constraints.Max, row, col)
		columns = append(columns,
			layout.Flexed(ColumnWeight, func(gtx c) d {
				return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx c) d {
					return f(gtx)
				})
			}),
		)
	}
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, columns...)
}

func colorBox(gtx c, size image.Point, r int, c int) d {
	defer op.Save(gtx.Ops).Load()

	clip.Rect{Max: size}.Add(gtx.Ops)
	switch gridInstance.Cells[c+(Columns*r)].CellType {
	case Empty:
		paint.ColorOp{Color: DefaultCellColor}.Add(gtx.Ops)
	case Wall:
		paint.ColorOp{Color: BlueCellColor}.Add(gtx.Ops)
	case Start:
		paint.ColorOp{Color: StartCellColor}.Add(gtx.Ops)
	case Finish:
		paint.ColorOp{Color: FinishCellColor}.Add(gtx.Ops)
	}
	paint.PaintOp{}.Add(gtx.Ops)

	// Confine the area of interest to a 100x100 rectangle.
	pointer.Rect(image.Rect(0, 0, 100, 100)).Add(gtx.Ops)
	// Declare the tag.
	pointer.InputOp{
		Tag:   gridInstance.Cells[c+(Columns*r)].Tag,
		Types: pointer.Press,
	}.Add(gtx.Ops)

	return layout.Dimensions{Size: size}
}
