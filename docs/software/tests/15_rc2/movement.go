package main

import (
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
			setServo("X", 180.0)
			currentX = 180
		case deltaX < 0:
			setServo("X", 0.0)
			currentX = 0
		default:
			setServo("X", float64(deltaX))
			currentX = deltaX
		}

	case "axisY":

		if invertY {
			angle = -angle
		}
		angle *= calibrateY

		switch deltaY := currentY + int(math.Round(angle)); {
		case deltaY > 180:
			setServo("Y", 180.0)
			currentY = 180
		case deltaY < 0:
			setServo("Y", 0.0)
			currentY = 0
		default:
			setServo("Y", float64(deltaY))
			currentY = deltaY
		}
	}
}

func calibrateServos() {
	log.Printf("Calibrating servomotors ...\n")
	centerServos()
	setServo("X", 0.0)
	setServo("Y", 0.0)

	setServo("X", 180.0)
	setServo("Y", 180.0)
	centerServos()
}

func centerServos() {
	log.Printf("Centering servos ...\n")
	setServo("X", 90.0)
	setServo("Y", 90.0)
	currentX = 90
	currentY = 90

}

func setServo(dir string, angle float64) {
	a := angle/900 + 0.05

	switch dir {
	case "X":
		servos.Apply(servoXpin, a)
	case "Y":
		servos.Apply(servoYpin, a)
	}
}
