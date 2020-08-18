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
		case newX > 150:
			setServo("X", 150)
			currentX = 150
		case newX < 30:
			setServo("X", 30)
			currentX = 30
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
		case newY > 150:
			setServo("Y", 150)
			currentY = 150
		case newY < 30:
			setServo("Y", 30)
			currentY = 30
		default:
			setServo("Y", newY)
			currentY = newY
		}
	}
}

// try both minimum and maximum of servos and then centers them
func calibrateServos() {
	log.Printf("Calibrating servomotors ...\n")
	setServo("X", 30)
	time.Sleep(380 * time.Millisecond)
	setServo("Y", 30)
	time.Sleep(380 * time.Millisecond)
	setServo("X", 150)
	time.Sleep(760 * time.Millisecond)
	setServo("Y", 150)
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
