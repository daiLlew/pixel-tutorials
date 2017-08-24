package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)


func main() {
	pixelgl.Run(run)
}

func run() {
	cgf := pixelgl.WindowConfig{
		Title: "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cgf)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Update()
	}
}