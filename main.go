package main

import (
	"image"
	"log"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
	"gocv.io/x/gocv"
)

// cams frames var img atomic.Value
var img atomic.Value

// idle
var idleTime = time.Now()
var idleStatus = false
var centered = false

// connection
var raspiAdaptor = raspi.NewAdaptor()

// list of devices
var camera = opencv.NewCameraDriver(cameraSource)
var servoX = gpio.NewServoDriver(raspiAdaptor, servoXpin)
var servoY = gpio.NewServoDriver(raspiAdaptor, servoYpin)

func main() {

	// work func
	work := func() {

		mat := gocv.NewMat()
		defer mat.Close()

		// store mat to img
		img.Store(mat)

		// turn camera on
		camera.On(opencv.Frame, func(data interface{}) {
			i := data.(gocv.Mat)
			// store i to img
			img.Store(i)
		})

		// calibration
		if calibration {
			calibrateServos()
		}

		// aplying max FPS
		if 1000/period > maxFPS {
			reducePeriod()
		}

		// loop
		gobot.Every(period*time.Millisecond, func() {

			//load an image
			i := img.Load().(gocv.Mat)
			if i.Empty() {
				return
			}

			// scan for objects
			objects := opencv.DetectObjects(cascade, i)

			// get a target's rectangle
			index, target := getNearestObject(objects)

			if (target != image.Rectangle{}) {

				// draw without target + target
				objectsNoTarget := append(objects[:index], objects[(index+1):]...)
				drawRects(i, []image.Rectangle{target}, targetColor)
				drawRects(i, objectsNoTarget, otherColor)

				// idle reset
				if idleStatus {
					idleStatus = false
				}

				// get a target's coordinate
				lock := getCoordinates(target)

				// aim if not in rectangle
				if !lock.In(midRect) {
					aimTarget(lock)
				}

			} else {

				// set new idleStatus
				if !idleStatus {
					idleTime = time.Now()
					idleStatus = true
					centered = false

				} else if !centered {

					// get the time difference
					if time.Now().Sub(idleTime).Seconds() >= idleDuration {
						centerServos()
						centered = true
						log.Printf("Centering servos, long time idle ...\n")
					}
				}
			}
		})
	}

	// variables
	connections := []gobot.Connection{raspiAdaptor}
	devices := []gobot.Device{camera, servoX, servoY}

	// set robot
	robot := gobot.NewRobot(
		robotName,
		connections,
		devices,
		work,
	)

	robot.Start()
}
