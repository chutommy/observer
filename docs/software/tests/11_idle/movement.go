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

		a *= invertX

		switch deltaX := currentX + a; {
		case deltaX > 180:
			_ = servoX.Move(uint8(180))
			currentX = 180
		case deltaX < 0:
			_ = servoX.Move(uint8(0))
			currentX = 0
		default:
			_ = servoX.Move(uint8(deltaX))
			currentX += a
		}

	case "axisY":

		a *= invertY

		switch deltaY := currentY + a; {
		case deltaY > 180:
			_ = servoY.Move(uint8(180))
			currentY = 180
		case deltaY < 0:
			_ = servoY.Move(uint8(0))
			currentY = 0
		default:
			_ = servoY.Move(uint8(currentY + a))
			currentY += a
		}
	}
}

func calibrateServos() {
	centerServos()
	_ = servoX.Min()
	_ = servoY.Min()
	_ = servoX.Max()
	_ = servoY.Max()
	centerServos()
}

func centerServos() {
	_ = servoX.Center()
	_ = servoY.Center()
}
