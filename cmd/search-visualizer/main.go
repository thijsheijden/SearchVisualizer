package main

import (
	"image/color"
	"log"
	"os"
	"search-visualizer/internal/grid"
	"search-visualizer/internal/menu"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

type (
	d = layout.Dimensions
	c = layout.Context
)

// Some NRGBA colors
var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
	cellColor  = color.NRGBA{R: 0xEC, G: 0xEC, B: 0xEC, A: 0xFF}
)

// UI variables
var (
	columnsInputTag = new(string)
	rowsInputTag    = new(string)
)

// Grid variables
var (
	gridRows         int     = 25
	gridColumns      int     = 25
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
	menu.New()
	grid.New()
	menu.PassCellTypeToGrid()
	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			menu.HandleInput(gtx)
			grid.HandleInput(gtx)
			// graphics.HandleCellClicks(gtx)
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Flexed(0.15, func(gtx c) d {
					return menu.Display(gtx)
				}),
				layout.Flexed(0.85, func(gtx c) d {
					return grid.Display(gtx)
				}),
			)
			e.Frame(gtx.Ops)
		}
	}

	return nil
}
