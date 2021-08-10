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

var camera = opencv.NewCameraDriver(cameraSource)
var servoX = gpio.NewServoDriver(raspiAdaptor, servoXpin)
var servoY = gpio.NewServoDriver(raspiAdaptor, servoYpin)

func main() {

	work := func() {

		mat := gocv.NewMat()
		defer func(mat *gocv.Mat) {
			_ = mat.Close()
		}(&mat)

		img.Store(mat)

		_ = camera.On(opencv.Frame, func(data interface{}) {
			i := data.(gocv.Mat)

			img.Store(i)
		})

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

			objects := opencv.DetectObjects(cascade, i)

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
		})
	}

	connections := []gobot.Connection{raspiAdaptor}
	devices := []gobot.Device{camera, servoX, servoY}

	robot := gobot.NewRobot(
		robotName,
		connections,
		devices,
		work,
	)

	_ = robot.Start()
}
