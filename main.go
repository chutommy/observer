package main

import (
	"fmt"
	"image/color"
	"log"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
	"gocv.io/x/gocv"
	"observer/engine"
	"observer/geometry"
)

const robotName = "Observer"

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

		// initialize servos
		servoX := engine.NewServo(servos, servoXpin, invertX, calibrateX, midPoint.X, tolerationX, pxsPerDegreeHor)
		servoY := engine.NewServo(servos, servoYpin, invertY, calibrateY, midPoint.Y, tolerationY, pxsPerDegreeVer)
		servoXY := engine.NewServos(servoX, servoY)

		// init center
		servoXY.CenterMiddleUp()
		time.Sleep(381 * time.Millisecond)

		// calibrate servos if enabled
		if calibration {
			servoXY.Calibrate()
		}

		// limit period according to maxFPS
		if 1000/period > maxFPS {
			reducePeriod()
		}

		time.Sleep(2 * time.Second)
		log.Println("Observing start running ...")

		// main loop every {period}ns
		gobot.Every(period*time.Millisecond, func() {

			//load the frame from img
			i := img.Load().(gocv.Mat)
			if i.Empty() {
				return
			}

			// scan for objects
			rects := opencv.DetectObjects(cascade1, i)
			objects := geometry.FromRects(rects)

			// if second cascade (optional)
			if cascade2 != "" {
				// objects = append(objects, opencv.DetectObjects(cascade2, i)...)
				rects = opencv.DetectObjects(cascade2, i)
				objects = append(objects, geometry.FromRects(rects)...)
			}

			// get the target's index and rectangle
			targetX := geometry.NearestObject(objects)
			target := objects[targetX]

			if targetX != -1 {

				// draw the target
				otherObjects := append(objects[:targetX], objects[(targetX+1):]...)
				target.Draw(&i, color.RGBA{
					R: uint8(targetColor.r),
					G: uint8(targetColor.g),
					B: uint8(targetColor.b),
					A: 0,
				}, targetColor.thickness)
				// draw non-target objects
				otherObjects.Draw(&i, color.RGBA{
					R: uint8(otherColor.r),
					G: uint8(otherColor.g),
					B: uint8(otherColor.b),
					A: 0,
				}, otherColor.thickness)

				// reset idle and suspend the counter
				if idleStatus {
					idleStatus = false
				}

				// get a target's coordinate
				lock := target.Center()

				// aim the target if it is not in the middle rectangle
				if !lock.In(midRect) {
					servoXY.Aim(lock)
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
						servoXY.Center()
						centered = true
					}
				}
			}

			// show window
			if windowed {

				// draw a mid rect
				geometry.FromRect(midRect).Draw(&i, color.RGBA{
					R: uint8(midRectColor.r),
					G: uint8(midRectColor.g),
					B: uint8(midRectColor.b),
					A: 0,
				}, midRectColor.thickness)

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
