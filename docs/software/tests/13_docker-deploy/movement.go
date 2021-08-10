package main

import (
	"fmt"
	"image"
	"log"
	"math"
)

var currentX = 90
var currentY = 90

func aimTarget(coor image.Point) {

	angleX := float64(coor.X-midPoint.X) / pxsPerDegree
	moveCam("axisX", angleX)

	angleY := float64(coor.Y-midPoint.Y) / pxsPerDegree
	moveCam("axisY", angleY)
}

func moveCam(direct string, angle float64) {

	switch direct {

	case "axisX":

		if invertX {
			angle = -angle
		}
		angle *= calibrateX

		switch deltaX := currentX + int(math.Round(angle)); {
		case deltaX > 180:
			_ = servoX.Max()

			currentX = 180
		case deltaX < 0:
			_ = servoY.Min()

			currentX = 0
		default:
			_ = servoX.Move(uint8(deltaX))
			currentX = deltaX
		}
		fmt.Println("XXX", currentX)

	case "axisY":

		if invertY {
			angle = -angle
		}
		angle *= calibrateY

		switch deltaY := currentY + int(math.Round(angle)); {
		case deltaY > 180:
			_ = servoY.Max()

			currentY = 180
		case deltaY < 0:
			_ = servoY.Min()

			currentY = 0
		default:
			_ = servoY.Move(uint8(deltaY))
			currentY = deltaY
		}

		fmt.Println("YYY", currentY)
	}
}

func calibrateServos() {
	log.Printf("Calibrating servomotors ...\n")
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
	currentX = 90
	currentY = 90
}
