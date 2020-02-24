package main

import (
	"image"
	"log"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
	"gocv.io/x/gocv"
)

// variable that cam stores its frames into
var img atomic.Value

// idle variables
var idleTime = time.Now()
var idleStatus = false
var centered = false

// raspberry pi adaptor (on boot)
var raspiAdaptor = raspi.NewAdaptor()

// a window variable
var window *opencv.WindowDriver

// camera driver
var camera = opencv.NewCameraDriver(cameraSource)

func main() {

	// only if window is enabled
	if windowed {
		window = opencv.NewWindowDriver()
	}

	// func which robot does
	work := func() {

		// creating a surface for frames
		mat := gocv.NewMat()
		defer mat.Close()
		img.Store(mat)

		// camera capturing
		camera.On(opencv.Frame, func(data interface{}) {
			i := data.(gocv.Mat)
			// store i to img
			img.Store(i)
		})

		// calibration at start if it is enabled
		if calibration {
			calibrateServos()
		}

		// limit period according to maxFPS
		if 1000/period > maxFPS {
			reducePeriod()
		}

		// main loop every {period}ns
		gobot.Every(period*time.Millisecond, func() {

			//load the frame from img
			i := img.Load().(gocv.Mat)
			if i.Empty() {
				return
			}

			// scan for objects using cascade
			objects := opencv.DetectObjects(cascade, i)

			// get a target's rectangle
			index, target := getNearestObject(objects)

			if (target != image.Rectangle{}) {

				// draw notarget objects + the target
				objectsNoTarget := append(objects[:index], objects[(index+1):]...)
				drawRects(i, []image.Rectangle{target}, targetColor)
				drawRects(i, objectsNoTarget, otherColor)

				// idle reset, suspend the counter
				if idleStatus {
					idleStatus = false
				}

				// get a target's coordinate
				lock := getCoordinates(target)

				// aim the target if it is not in the middle rectangle
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

					// get the time difference, if idle too long - reset
					if time.Now().Sub(idleTime).Seconds() >= idleDuration {
						centerServos()
						centered = true
						log.Printf("Centering servos, long time idle ...\n")
					}
				}
			}

			// show window
			if windowed {
				drawRects(i, []image.Rectangle{midRect}, midRectColor)
				window.ShowImage(i)
				window.WaitKey(1)
			}

		})
	}

	// define adaptors and devices
	connections := []gobot.Connection{raspiAdaptor}
	devices := []gobot.Device{camera, servoX, servoY}

	// adds window if window is enabled
	if windowed {
		devices = append(devices, window)
	}

	// set robot atributes
	robot := gobot.NewRobot(
		robotName,
		connections,
		devices,
		work,
	)

	robot.Start()
}
