package main

import (
	"image"

	"gobot.io/x/gobot/platforms/opencv"
	"gocv.io/x/gocv"
)

// color type
type cusColor struct {
	r, g, b, thickness int
}

// draw rectangle rect into image img with color type col
func drawRects(img gocv.Mat, rects []image.Rectangle, col cusColor) {
	opencv.DrawRectangles(img, rects, col.r, col.g, col.b, col.thickness)
}
