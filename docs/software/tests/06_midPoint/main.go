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

var robotName = "Big Brother"
var cascade = "../haarcascades/haarcascade_frontalface_default.xml"
var cameraSource = 0

var camWidth = 640
var camHeight = 480

var img atomic.Value

var midPoint = image.Point{
	X: camWidth / 2,
	Y: camHeight / 2,
}

func main() {
	r := raspi.NewAdaptor()

	window := opencv.NewWindowDriver()
	camera := opencv.NewCameraDriver(cameraSource)

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

		gobot.Every(1*time.Millisecond, func() {

			i := img.Load().(gocv.Mat)
			if i.Empty() {
				return
			}

			objects := opencv.DetectObjects(cascade, i)

			target := getNearestObject(objects)
			if (target != image.Rectangle{}) {

				lock := getCoordinates(target)

				opencv.DrawRectangles(i, []image.Rectangle{target, image.Rect(
					midPoint.X-5,
					midPoint.Y-5,
					midPoint.X+5,
					midPoint.Y+5,
				)}, 0, 255, 0, 5)
				fmt.Printf("X: %v\tY: %v\n", lock.X, lock.Y)
				window.ShowImage(i)
				window.WaitKey(1)
			}
		})
	}

	connections := []gobot.Connection{r}
	devices := []gobot.Device{camera, window}

	robot := gobot.NewRobot(
		robotName,
		connections,
		devices,
		work,
	)

	_ = robot.Start()
}
