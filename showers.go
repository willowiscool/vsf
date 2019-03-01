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
		if changed[i] {
			rect.Color = color.RGBA{CONFIG.CHANGED[0], CONFIG.CHANGED[1], CONFIG.CHANGED[2], CONFIG.CHANGED[3]}
		} else {
			rect.Color = color.RGBA{CONFIG.FG[0], CONFIG.FG[1], CONFIG.FG[2], CONFIG.FG[3]}
		}
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
		if changed[i] {
			points.Color = color.RGBA{CONFIG.CHANGED[0], CONFIG.CHANGED[1], CONFIG.CHANGED[2], CONFIG.CHANGED[3]}
		} else {
			points.Color = color.RGBA{CONFIG.FG[0], CONFIG.FG[1], CONFIG.FG[2], CONFIG.FG[3]}
		}
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
		if changed[i] {
			circle.Color = color.RGBA{CONFIG.CHANGED[0], CONFIG.CHANGED[1], CONFIG.CHANGED[2], CONFIG.CHANGED[3]}
		} else {
			circle.Color = color.RGBA{CONFIG.FG[0], CONFIG.FG[1], CONFIG.FG[2], CONFIG.FG[3]}
		}
		circle.Push(pixel.V(float64(CONFIG.BLOCK_HEIGHT_MULT * CONFIG.LIST_LENGTH), float64(CONFIG.BLOCK_HEIGHT_MULT * CONFIG.LIST_LENGTH)))
		circle.CircleArc(float64(val * CONFIG.BLOCK_HEIGHT_MULT), float64(i) * arcWidth, float64(i+1) * arcWidth, 0)
	}
	return circle
}
