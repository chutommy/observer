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

var robotName = "Big Brother"
var cascade = "../haarcascades/haarcascade_frontalface_default.xml"

var cameraSource = 0
var servoXpin = "servoMotX"
var servoYpin = "servoMotY"

var camWidth = 640
var camHeight = 480

var angleOfView = 72.4
var maxFPS time.Duration = 30

var aimArea = 0.25
var invertX = 1
var invertY = 1

var calibration = true

var period time.Duration = 35
var idleDuration float64 = 10

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
		log.Printf("Creating new Mat ...\n")

		img.Store(mat)

		_ = camera.On(opencv.Frame, func(data interface{}) {
			i := data.(gocv.Mat)

			img.Store(i)
		})
		log.Printf("First camera device initialization ...\n")

		if calibration {
			calibrateServos()
			log.Printf("Calibrating servomotors ...\n")
		}

		if 1000/period > maxFPS {
			log.Printf("Reducing period from %v to %v (max. FPS) ...", period, maxFPS)
			period = 1000 / maxFPS
		}

		gobot.Every(period*time.Millisecond, func() {

			i := img.Load().(gocv.Mat)
			if i.Empty() {
				return
			}

			objects := opencv.DetectObjects(cascade, i)

			target := getNearestObject(objects)
			if (target != image.Rectangle{}) {

				if idleStatus {
					idleStatus = false
				}

				lock := getCoordinates(target)

				if !lock.In(midRect) {
					aimTarget(lock)
				}

				opencv.DrawRectangles(i, []image.Rectangle{midRect}, 0, 255, 0, 5)
				window.ShowImage(i)
				window.WaitKey(1)

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
	devices := []gobot.Device{camera, window, servoX, servoY}

	robot := gobot.NewRobot(
		robotName,
		connections,
		devices,
		work,
	)

	_ = robot.Start()
}
