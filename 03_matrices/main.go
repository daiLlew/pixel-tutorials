package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"os"
	_ "image/png"
	"image"
	"path/filepath"
	"flag"
	"time"
)

var (
	imgPath        string
)

func main() {
	p := flag.String("imgPath", "", "File path of the png image to use for the example.")
	flag.Parse()

	imgPath = *p

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
	win.SetSmooth(true)

	path, _ := filepath.Abs(imgPath)
	pic, err := loadPicture(path)
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Firebrick)

	angle := 0.0
	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Firebrick)
		angle += 5 * dt

		mat := pixel.IM
		mat = mat.Rotated(pixel.ZV, angle)
		mat = mat.Moved(win.Bounds().Center())
		sprite.Draw(win, mat)

		win.Update()
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
