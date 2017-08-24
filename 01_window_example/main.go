package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cgf := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true, // VSync aligns the window refresh with your monitor refresh.
	}

	win, err := pixelgl.NewWindow(cgf)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Blueviolet)

	for !win.Closed() {
		win.Update()
	}
}
