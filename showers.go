package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"image/color"
	"math"
)

func rectDraw(list []int, changed []bool) *imdraw.IMDraw {
	rect := imdraw.New(nil)
	for i, val := range list {
		setColor(float64(val), changed[i], rect)
		rect.Push(pixel.V(
			float64(i * CONFIG.BLOCK_WIDTH),
			float64(val * CONFIG.BLOCK_HEIGHT_MULT)))
		rect.Push(pixel.V(float64((i+1) * CONFIG.BLOCK_WIDTH), 0))
		rect.Rectangle(0)
	}
	return rect
}

func pointDraw(list []int, changed []bool) *imdraw.IMDraw {
	points := imdraw.New(nil)
	for i, val := range list {
		setColor(float64(val), changed[i], points)
		points.Push(pixel.V(
			float64(i * CONFIG.BLOCK_WIDTH),
			float64(val * CONFIG.BLOCK_HEIGHT_MULT - CONFIG.BLOCK_HEIGHT_MULT)))
		points.Push(pixel.V(float64((i+1) * CONFIG.BLOCK_WIDTH), float64(val * CONFIG.BLOCK_HEIGHT_MULT)))
		points.Rectangle(0)
	}
	return points
}

func circleDraw(list []int, changed []bool) *imdraw.IMDraw {
	arcWidth := 2*math.Pi / float64(CONFIG.LIST_LENGTH)
	circle := imdraw.New(nil)
	for i, val := range list {
		setColor(float64(val), changed[i], circle)
		circle.Push(pixel.V(float64(CONFIG.BLOCK_HEIGHT_MULT * CONFIG.LIST_LENGTH), float64(CONFIG.BLOCK_HEIGHT_MULT * CONFIG.LIST_LENGTH)))
		circle.CircleArc(float64(val * CONFIG.BLOCK_HEIGHT_MULT), float64(i) * arcWidth, float64(i+1) * arcWidth, 0)
	}
	return circle
}

func setColor(val float64, itchanged bool, drawer *imdraw.IMDraw) {
	if !CONFIG.RAINBOW {
		if itchanged {
			drawer.Color = color.RGBA{CONFIG.CHANGED[0], CONFIG.CHANGED[1], CONFIG.CHANGED[2], CONFIG.CHANGED[3]}
		} else {
			drawer.Color = color.RGBA{CONFIG.FG[0], CONFIG.FG[1], CONFIG.FG[2], CONFIG.FG[3]}
		}
	} else {
		r := uint8(math.Floor(math.Sin(math.Pi / float64(CONFIG.LIST_LENGTH) * 2 * val) * 127) + 128)
		g := uint8(math.Floor(math.Sin(math.Pi / float64(CONFIG.LIST_LENGTH) * 2 * val + math.Pi * 2/3) * 127) + 128)
		b := uint8(math.Floor(math.Sin(math.Pi / float64(CONFIG.LIST_LENGTH) * 2 * val + 2 * math.Pi * 2/3) * 127) + 128)
		drawer.Color = color.RGBA{r, g, b, 1}
	}
}
