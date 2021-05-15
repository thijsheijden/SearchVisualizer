package menu

import (
	"image"
	"image/color"
	"search-visualizer/internal/grid"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// Creating new types to reduce the clutter
type (
	c = layout.Context
	d = layout.Dimensions
)

// Display displays the menu
func Display(gtx c) d {
	// Uniformly inset the entire menu by 8 dp
	return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx c) d {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(0.4, func(gtx c) d {
				return topMenu.displayAllInputsWithLabels(gtx)
			}),
			layout.Flexed(0.4, func(gtx c) d {
				return topMenu.displayCellPaintSquare(gtx)
			}),
			layout.Flexed(0.2, func(gtx c) d {
				return topMenu.displayStartStopButtons(gtx)
			}),
		)

	})
}

func (m *menu) displayAllInputsWithLabels(gtx c) d {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx c) d {
			return m.displayInputWithLabel(gtx, m.gridColumnsInputLabel, "Columns", m.gridColumnsInput)
		}),
		layout.Flexed(1, func(gtx c) d {
			return m.displayInputWithLabel(gtx, m.gridRowsInputLabel, "Rows", m.gridRowsInput)
		}),
		layout.Flexed(1, func(gtx c) d {
			return m.displayInputWithLabel(gtx, m.tickSpeedInputLabel, "Tick speed", m.tickSpeedInput)
		}),
	)
}

func (m *menu) displayInputWithLabel(gtx c, label *widget.Label, labelText string, input *widget.Editor) d {
	return layout.Inset{Top: unit.Dp(2), Bottom: unit.Dp(2)}.Layout(gtx, func(gtx c) d {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(1, func(gtx c) d {
				return m.displayLabel(gtx, label, labelText, false)
			}),
			layout.Flexed(1, func(gtx c) d {
				return m.displayInputField(gtx, input)
			}),
		)
	})
}

func (m *menu) displayInputField(gtx c, input *widget.Editor) d {
	e := material.Editor(m.theme, input, "25")
	e.Editor.SingleLine = true
	e.TextSize = unit.Dp(14)
	border := widget.Border{
		Color:        color.NRGBA{R: 0x77, G: 0xDD, B: 0x77, A: 0xFF},
		CornerRadius: unit.Dp(4),
		Width:        unit.Dp(1),
	}
	return border.Layout(gtx, func(gtx c) d {
		return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx c) d {
			return e.Layout(gtx)
		})
	})
}

func (m *menu) displayLabel(gtx c, label *widget.Label, labelText string, centered bool) d {
	l := material.Label(m.theme, unit.Dp(14), labelText)
	if centered {
		l.Alignment = text.Middle
	}
	return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx c) d {
		return l.Layout(gtx)
	})
}

func (m *menu) displayCellPaintSquare(gtx c) d {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(0.8, func(gtx c) d {
			return layout.Center.Layout(gtx, func(gtx c) d {
				return m.colorBox(gtx, gtx.Constraints.Max)
			})
		}),
		layout.Flexed(0.2, func(gtx c) d {
			return m.displayLabel(gtx, m.cellPaintTypeLabel, m.cellPaintType.String(), true)
		}),
	)
}

func (m *menu) colorBox(gtx c, size image.Point) d {
	defer op.Save(gtx.Ops).Load()

	// Make it a square
	var squareSize image.Point
	if size.X < size.Y {
		squareSize = image.Point{X: size.X, Y: size.X}
	} else {
		squareSize = image.Point{X: size.Y, Y: size.Y}
	}

	clip.Rect{Max: squareSize}.Add(gtx.Ops)
	switch m.cellPaintType {
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
	pointer.Rect(image.Rect(0, 0, squareSize.X, squareSize.Y)).Add(gtx.Ops)
	// Declare the tag.
	pointer.InputOp{
		Tag:   m.cellPaintTag,
		Types: pointer.Scroll,
	}.Add(gtx.Ops)

	return layout.Dimensions{Size: squareSize}
}

func (m *menu) displayStartStopButtons(gtx c) d {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx c) d {
			return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx, func(gtx c) d {
				return m.displayButton(gtx, m.startButton, "Start", grid.FinishCellColor)
			})
		}),
		layout.Flexed(1, func(gtx c) d {
			return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx, func(gtx c) d {
				return m.displayButton(gtx, m.stopButton, "Stop", grid.StartCellColor)
			})
		}),
	)
}

func (m *menu) displayButton(gtx c, button *widget.Clickable, buttonText string, color color.NRGBA) d {
	b := material.Button(m.theme, button, buttonText)
	b.Background = color
	b.TextSize = unit.Dp(14)
	return b.Layout(gtx)
}
