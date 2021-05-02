package graphics

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func DisplayGrid(gtx C) D {
	var rows []layout.FlexChild
	for i := 0; i < gridRows; i++ {
		rows = append(rows,
			layout.Flexed(gridRowWeight, func(gtx C) D {
				return drawRow(gtx)
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

func drawRow(gtx C) D {
	var columns []layout.FlexChild
	for i := 0; i < gridColumns; i++ {
		columns = append(columns,
			layout.Flexed(gridColumnWeight, func(gtx C) D {
				return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx C) D {
					return colorBox(gtx, gtx.Constraints.Max, cellColor)
				})
			}),
		)
	}
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, columns...)
}

func colorBox(gtx C, size image.Point, color color.NRGBA) D {
	defer op.Save(gtx.Ops).Load()
	var newSize image.Point
	if size.X < size.Y {
		newSize = image.Point{
			X: size.X,
			Y: size.X,
		}
	} else {
		newSize = image.Point{
			X: size.Y,
			Y: size.Y,
		}
	}
	clip.Rect{Max: newSize}.Add(gtx.Ops)
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: newSize}
}
