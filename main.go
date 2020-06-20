package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
)

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
	work := botWork(camera, window)

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
