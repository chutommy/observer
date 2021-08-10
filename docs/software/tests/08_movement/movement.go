package main

import (
	"image"
	"math"
)

var currentX = 90
var currentY = 90

var pixelsX = float64(camWidth) / angleOfView
var pixelsY = float64(camHeight) / angleOfView

func aimTarget(coor image.Point) {

	angleX := float64(coor.X-midPoint.X) / pixelsX
	moveCam("axisX", angleX)

	angleY := float64(coor.Y-midPoint.Y) / pixelsY
	moveCam("axisY", angleY)
}

func moveCam(direct string, angle float64) {
	a := int(math.Round(angle))

	switch direct {
	case "axisX":
		_ = servoX.Move(uint8(currentX + a))
	case "axisY":
		_ = servoY.Move(uint8(currentY + a))
	}
}
