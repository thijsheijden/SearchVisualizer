package menu

import (
	"image/color"

	"gioui.org/layout"
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
		return topMenu.displayAllInputsWithLabels(gtx)
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
	)
}

func (m *menu) displayInputWithLabel(gtx c, label *widget.Label, labelText string, input *widget.Editor) d {
	return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx, func(gtx c) d {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(1, func(gtx c) d {
				return m.displayLabel(gtx, label, labelText)
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

func (m *menu) displayLabel(gtx c, label *widget.Label, text string) d {
	l := material.Label(m.theme, unit.Dp(14), text)
	return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx c) d {
		return l.Layout(gtx)
	})
}
