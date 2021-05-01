package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

// Some NRGBA colors
var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

// Grid variables
var (
	gridRows         int     = 10
	gridColumns      int     = 10
	gridRowWeight    float32 = 1 / float32(gridRows)
	gridColumnWeight float32 = 1 / float32(gridColumns)
)

func main() {
	go func() {
		w := app.NewWindow(app.Size(unit.Px(1024), unit.Px(800)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	var ops op.Ops
	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			grid(gtx)
			e.Frame(gtx.Ops)
		}
	}

	return nil
}

func colorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer op.Save(gtx.Ops).Load()
	clip.Rect{Max: size}.Add(gtx.Ops)
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func drawRow(gtx layout.Context) layout.Dimensions {
	var columns []layout.FlexChild
	for i := 0; i < gridColumns; i++ {
		columns = append(columns,
			layout.Flexed(gridColumnWeight, func(gtx layout.Context) layout.Dimensions {
				return inset(gtx)
			}),
		)
	}
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, columns...)
}

func inset(gtx layout.Context) layout.Dimensions {
	// Draw rectangle inset by 8dp on every side
	return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return colorBox(gtx, gtx.Constraints.Max, red)
	})
}

func grid(gtx layout.Context) layout.Dimensions {
	var rows []layout.FlexChild
	for i := 0; i < gridColumns; i++ {
		rows = append(rows,
			layout.Flexed(gridRowWeight, func(gtx layout.Context) layout.Dimensions {
				return drawRow(gtx)
			}),
		)
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rows...)
}
