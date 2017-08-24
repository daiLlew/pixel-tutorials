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
)

var imgPath string

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

	path, _ := filepath.Abs(imgPath)
	pic, err := loadPicture(path)
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Lightblue)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
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
