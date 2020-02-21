package main

import (
	"image"

	"gobot.io/x/gobot/platforms/opencv"
	"gocv.io/x/gocv"
)

// type for custom color
type cusColor struct {
	r, g, b, thickness int
}

// drawRects draw rects with a cusColor
func drawRects(img gocv.Mat, rects []image.Rectangle, col cusColor) {
	opencv.DrawRectangles(img, rects, col.r, col.g, col.b, col.thickness)
}
