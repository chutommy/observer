package main

import (
	"image"
	"log"
	"math"
	"time"
)

// current servos states
var currentX = 90
var currentY = 90

// aim target by its coordinates
func aimTarget(coor image.Point) {

	xDiff := float64(coor.X - midPoint.X)
	yDiff := float64(coor.Y - midPoint.Y)

	// equalize X axis
	if (xDiff > tolerationX) || (xDiff < tolerationXr) {
		angleX := xDiff / pxsPerDegreeHor
		moveCam("axisX", angleX)
	}

	// equalize Y axis
	if (yDiff > tolerationY) || (yDiff < tolerationYr) {
		angleY := yDiff / pxsPerDegreeVer
		moveCam("axisY", angleY)
	}
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
		switch newX := currentX + int(math.Round(angle)); {
		case newX > 180:
			setServo("X", 180)
			currentX = 180
		case newX < 0:
			setServo("X", 0)
			currentX = 0
		default:
			setServo("X", newX)
			currentX = newX
		}

	// Y movement
	case "axisY":

		// Y invert/calibrate
		if invertY {
			angle = -angle
		}
		angle *= calibrateY

		// movement selection Y
		switch newY := currentY + int(math.Round(angle)); {
		case newY > 180:
			setServo("Y", 180)
			currentY = 180
		case newY < 0:
			setServo("Y", 0)
			currentY = 0
		default:
			setServo("Y", newY)
			currentY = newY
		}
	}
}

// try both minimum and maximum of servos and then centers them
func calibrateServos() {
	log.Printf("Calibrating servomotors ...\n")
	setServo("X", 0)
	time.Sleep(380 * time.Millisecond)
	setServo("Y", 0)
	time.Sleep(380 * time.Millisecond)
	setServo("X", 180)
	time.Sleep(760 * time.Millisecond)
	setServo("Y", 180)
	time.Sleep(760 * time.Millisecond)
	centerServos()
	time.Sleep(380 * time.Millisecond)
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
