package main

import (
	"image"
	"log"
	"math"
)

// current servos states
var currentX = 90
var currentY = 90

// aim target by its coordinates
func aimTarget(coor image.Point) {

	// equalize X axis
	angleX := float64(coor.X-midPoint.X) / pxsPerDegree
	moveCam("axisX", angleX)

	// equalize Y axis
	angleY := float64(coor.Y-midPoint.Y) / pxsPerDegree
	moveCam("axisY", angleY)
}

// move cam by angle
func moveCam(direct string, angle float64) {

	// choose a direction
	switch direct {

	// X movement
	case "axisX":

		// X invert/calibrate
		if invertX {
			angle = -angle
		}
		angle *= calibrateX

		// movement selection X
		switch deltaX := currentX + int(math.Round(angle)); {
		case deltaX > 180:
			setServo("X", 180)
			currentX = 180
		case deltaX < 0:
			setServo("X", 0)
			currentX = 0
		default:
			setServo("X", deltaX)
			currentX = deltaX
		}

	// Y movement
	case "axisY":

		// Y invert/calibrate
		if invertY {
			angle = -angle
		}
		angle *= calibrateY

		// movement selection Y
		switch deltaY := currentY + int(math.Round(angle)); {
		case deltaY > 180:
			setServo("Y", 180)
			currentY = 180
		case deltaY < 0:
			setServo("Y", 0)
			currentY = 0
		default:
			setServo("Y", deltaY)
			currentY = deltaY
		}
	}
}

// try both minimum and maximum of servos and then centers them
func calibrateServos() {
	log.Printf("Calibrating servomotors ...\n")
	centerServos()
	setServo("X", 0)
	setServo("Y", 0)
	// WAIT
	setServo("X", 180)
	setServo("Y", 180)
	centerServos()
}

// set servos to the default postion
func centerServos() {
	log.Printf("Centering servos ...\n")
	setServo("X", 90)
	setServo("Y", 90)
	currentX = 90
	currentY = 90
	// WAIT
}

// move the servo to the angle
func setServo(direct string, angle int) {
	a := float64(angle)/900 + 0.05

	switch direct {
	case "X":
		servos.Apply(servoXpin, a)
	case "Y":
		servos.Apply(servoYpin, a)
	}
}
