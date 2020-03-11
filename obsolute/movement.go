package main

import (
	"image"
	"log"
	"math"
)

// current servos statuses
var currentX = 90
var currentY = 90

// aimTarget aims target by its coordinates
func aimTarget(coor image.Point) {

	// equalize X axis
	angleX := float64(coor.X-midPoint.X) / pxsPerDegree
	moveCam("axisX", angleX)

	// equalize Y axis
	angleY := float64(coor.Y-midPoint.Y) / pxsPerDegree
	moveCam("axisY", angleY)
}

// moveCam moves cam on X or Y axis
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

		// movement selection
		switch deltaX := currentX + int(math.Round(angle)); {
		case deltaX > 180:
			servoX.Max()
			//servoX.Move(uint8(180))
			currentX = 180
		case deltaX < 0:
			servoY.Min()
			//servoX.Move(uint8(0))
			currentX = 0
		default:
			servoX.Move(uint8(deltaX))
			currentX = deltaX
		}

	// Y movement
	case "axisY":

		// Y invert/calibrate
		if invertY {
			angle = -angle
		}
		angle *= calibrateY

		// movement selection
		switch deltaY := currentY + int(math.Round(angle)); {
		case deltaY > 180:
			servoY.Max()
			//servoY.Move(uint8(180))
			currentY = 180
		case deltaY < 0:
			servoY.Min()
			//servoY.Move(uint8(0))
			currentY = 0
		default:
			servoY.Move(uint8(deltaY))
			currentY = deltaY
		}
	}
}

// calibrateServos tries both minimum and maximum of servos and then certers them
func calibrateServos() {
	log.Printf("Calibrating servomotors ...\n")
	centerServos()
	servoX.Min()
	servoY.Min()
	servoX.Max()
	servoY.Max()
	centerServos()
}

// centerServos set servos to the default postion
func centerServos() {
	servoX.Center()
	servoY.Center()
	currentX = 90
	currentY = 90
}
