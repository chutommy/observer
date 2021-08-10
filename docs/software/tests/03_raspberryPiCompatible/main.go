package main

import (
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

var img atomic.Value

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

		gobot.Every(10*time.Millisecond, func() {
			i := img.Load().(gocv.Mat)
			if i.Empty() {
				return
			}
			faces := opencv.DetectObjects(cascade, i)
			opencv.DrawRectangles(i, faces, 0, 255, 0, 5)
			window.ShowImage(i)
			window.WaitKey(1)
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
