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

var img atomic.Value

var idleTime = time.Now()
var idleStatus = false
var centered = false

var raspiAdaptor = raspi.NewAdaptor()

var window = opencv.NewWindowDriver()
var camera = opencv.NewCameraDriver(cameraSource)
var servoX = gpio.NewServoDriver(raspiAdaptor, servoXpin)
var servoY = gpio.NewServoDriver(raspiAdaptor, servoYpin)

func main() {

	work := func() {

		mat := gocv.NewMat()
		defer func(mat *gocv.Mat) {
			_ = mat.Close()
		}(&mat)
		log.Printf("Initializating new Mat ...\n")

		img.Store(mat)

		_ = camera.On(opencv.Frame, func(data interface{}) {
			i := data.(gocv.Mat)

			img.Store(i)
		})
		log.Printf("Initializating camera device ...\n")

		if calibration {
			calibrateServos()
		}

		if 1000/period > maxFPS {
			reducePeriod()
		}

		gobot.Every(period*time.Millisecond, func() {

			i := img.Load().(gocv.Mat)
			if i.Empty() {
				return
			}

			objects := sumCascades(i)

			index, target := getNearestObject(objects)

			if (target != image.Rectangle{}) {

				objectsNoTarget := append(objects[:index], objects[(index+1):]...)
				drawRects(i, []image.Rectangle{target}, targetColor)
				drawRects(i, objectsNoTarget, otherColor)

				if idleStatus {
					idleStatus = false
				}

				lock := getCoordinates(target)

				if !lock.In(midRect) {
					aimTarget(lock)
				}

			} else {

				if !idleStatus {
					idleTime = time.Now()
					idleStatus = true
					centered = false

				} else if !centered {

					if time.Now().Sub(idleTime).Seconds() >= idleDuration {
						centerServos()
						centered = true
						log.Printf("Centering servos, long time idle ...\n")
					}
				}
			}

			drawRects(i, []image.Rectangle{midRect}, midRectColor)
			window.ShowImage(i)
			window.WaitKey(1)

		})
	}

	connections := []gobot.Connection{raspiAdaptor}
	devices := []gobot.Device{window, camera, servoX, servoY}

	robot := gobot.NewRobot(
		robotName,
		connections,
		devices,
		work,
	)

	_ = robot.Start()
}
