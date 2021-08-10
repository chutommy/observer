package main

import (
	"image"

	"gobot.io/x/gobot/platforms/opencv"
	"gocv.io/x/gocv"
)

type cusColor struct {
	r, g, b, thickness int
}

func drawRects(img gocv.Mat, rects []image.Rectangle, col cusColor) {
	opencv.DrawRectangles(img, rects, col.r, col.g, col.b, col.thickness)
}
