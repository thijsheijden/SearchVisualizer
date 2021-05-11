package graphics

import (
	"image"
	"search-visualizer/internal/grid"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func DisplayGrid(gtx C) D {
	var rows []layout.FlexChild
	for r := 0; r < grid.Rows; r++ {
		f := drawRowFunc(gtx, r)
		rows = append(rows,
			layout.Flexed(grid.RowWeight, func(gtx C) D {
				return f(gtx)
			}),
		)
	}
	return layout.Inset{
		Top:    unit.Dp(0),
		Left:   unit.Dp(8),
		Right:  unit.Dp(8),
		Bottom: unit.Dp(8),
	}.Layout(gtx, func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical, Alignment: layout.Alignment(layout.Center)}.Layout(gtx, rows...)
	})
}

func drawRowFunc(gtx C, r int) func(C) D {
	return func(gtx C) D {
		return drawRow(gtx, r)
	}
}

func drawBoxFunc(gtx C, size image.Point, r int, c int) func(C) D {
	return func(gtx C) D {
		return colorBox(gtx, gtx.Constraints.Max, r, c)
	}
}

func drawRow(gtx C, r int) D {
	var columns []layout.FlexChild
	for c := 0; c < grid.Columns; c++ {
		f := drawBoxFunc(gtx, gtx.Constraints.Max, r, c)
		columns = append(columns,
			layout.Flexed(grid.ColumnWeight, func(gtx C) D {
				return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx C) D {
					return f(gtx)
				})
			}),
		)
	}
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, columns...)
}

func colorBox(gtx C, size image.Point, r int, c int) D {
	defer op.Save(gtx.Ops).Load()
	// var newSize image.Point
	// if size.X < size.Y {
	// 	newSize = image.Point{
	// 		X: size.X,
	// 		Y: size.X,
	// 	}
	// } else {
	// 	newSize = image.Point{
	// 		X: size.Y,
	// 		Y: size.Y,
	// 	}
	// }

	clip.Rect{Max: size}.Add(gtx.Ops)
	switch grid.GridInstance.Cells[c+(grid.Columns*r)].CellType {
	case grid.Empty:
		paint.ColorOp{Color: grid.DefaultCellColor}.Add(gtx.Ops)
	case grid.Wall:
		paint.ColorOp{Color: grid.BlueCellColor}.Add(gtx.Ops)
	case grid.Start:
		paint.ColorOp{Color: grid.StartCellColor}.Add(gtx.Ops)
	case grid.Finish:
		paint.ColorOp{Color: grid.FinishCellColor}.Add(gtx.Ops)
	}
	paint.PaintOp{}.Add(gtx.Ops)

	// Confine the area of interest to a 100x100 rectangle.
	pointer.Rect(image.Rect(0, 0, 100, 100)).Add(gtx.Ops)
	// Declare the tag.
	pointer.InputOp{
		Tag:   grid.GridInstance.Cells[c+(grid.Columns*r)].Tag,
		Types: pointer.Press,
	}.Add(gtx.Ops)

	return layout.Dimensions{Size: size}
}
