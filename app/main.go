package main

import (
	"fmt"
	"image"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
	"gocv.io/x/gocv"
)

// frames
var img atomic.Value

// idle variables
var idleTime = time.Now()
var idleStatus = false
var centered = false

// raspberry pi adaptor
var raspiAdaptor = raspi.NewAdaptor()

// showing camera's view
var window *opencv.WindowDriver

// camera driver
var camera = opencv.NewCameraDriver(cameraSource)

func main() {

	// enable window driver
	if windowed {
		window = opencv.NewWindowDriver()
	}

	// robot's func
	work := func() {

		// creating a storage for frames
		mat := gocv.NewMat()
		defer mat.Close()
		img.Store(mat)

		// turn camera on
		camera.On(opencv.Frame, func(data interface{}) {
			i := data.(gocv.Mat)
			// store the frame
			img.Store(i)
		})

		// init center
		centerServos()
		time.Sleep(381 * time.Millisecond)

		// calibrate servos if enabled
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

			// scan for objects
			objects := opencv.DetectObjects(cascade1, i)

			// if second cascade (optional)
			if cascade2 != "" {
				objects = append(objects, opencv.DetectObjects(cascade2, i)...)
			}

			// get the target's index and rectangle
			index, target := getNearestObject(objects)

			if index != -1 {

				// draw notarget objects + the target
				objectsNoTarget := append(objects[:index], objects[(index+1):]...)
				drawRects(i, []image.Rectangle{target}, targetColor)
				drawRects(i, objectsNoTarget, otherColor)

				// reset idle and suspend the counter
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
						fmt.Println("Idle to long ...")
						centerServos()
						centered = true
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
	devices := []gobot.Device{camera}

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
