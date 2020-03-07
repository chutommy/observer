package main

import (
	"image"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
	"gocv.io/x/gocv"
)

var img atomic.Value

func main() {

	raspiAdaptor := raspi.NewAdaptor()
	camera := opencv.NewCameraDriver(0)
	window := opencv.NewWindowDriver()

	work := func() {

		mat := gocv.NewMat()
		defer mat.Close()
		img.Store(mat)

		camera.On(opencv.Frame, func(data interface{}) {
			i := data.(gocv.Mat)
			img.Store(i)
		})

		gobot.Every(100*time.Millisecond, func() {
			i := img.Load().(gocv.Mat)
			if i.Empty() {
				return
			}

			opencv.DrawRectangles(i, []image.Rectangle{image.Rectangle{image.Point{10, 10}, image.Point{600, 400}}}, 200, 20, 20, 2)

			window.ShowImage(i)
			window.WaitKey(1)
		})
	}

	robot := gobot.NewRobot(
		"r",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{camera, window},
		work,
	)

	robot.Start()
}
