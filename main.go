package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/imdraw"

	"math/rand"
	"image/color"
)
const (
	LIST_LENGTH = 100 //the length of the list to be sorted
	BLOCK_WIDTH = 10 //the width of each block
	BLOCK_HEIGHT_MULT = 5 //the amount that the height of the block is multiplied by its position
	WIDTH = LIST_LENGTH * BLOCK_WIDTH //the width of the window
	HEIGHT = LIST_LENGTH * BLOCK_HEIGHT_MULT //the height of the window
)
var (
	imd *imdraw.IMDraw
	list = rand.Perm(LIST_LENGTH)
)

func main() {
	for i, _ := range list {
		list[i]++
	}
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	
	//background
	imd = imdraw.New(nil)
	imd.Color = pixel.RGB(0, 1, 0)
	imd.Push(pixel.V(0, 0))
	imd.Color = pixel.RGB(0, 1, 1)
	imd.Push(pixel.V(WIDTH, HEIGHT))
	imd.Rectangle(0)


	for !win.Closed() {
		win.Clear(color.RGBA{0x55, 0xaa, 0xaa, 0xff})
		imd.Draw(win)
		for i, val := range list {
			rect := imdraw.New(nil)
			rect.Color = pixel.RGB(1, 1, 1)
			rect.Push(pixel.V(
				float64(i * BLOCK_WIDTH),
				float64(val * BLOCK_HEIGHT_MULT)))
			rect.Push(pixel.V(float64((i+1) * BLOCK_WIDTH), 0))
			rect.Rectangle(0)
			rect.Draw(win)
		}
		win.Update()
	}
}
