package main

import (
	"fmt"
	"image"
	"log"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
	"gocv.io/x/gocv"
)

// frames
var img atomic.Value

// idle variables
var idleTime = time.Now()
var idleStatus = false
var centered = false

func botWork(camera *opencv.CameraDriver, window *opencv.WindowDriver) func() {

	return func() {

		// creating a storage for frames
		ma := gocv.NewMat()
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
}
