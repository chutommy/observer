package main

import (
	"errors"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
	"gocv.io/x/gocv"
	"observer/config"
	"observer/engine"
	"observer/geometry"
	"observer/observerconfig"
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

	// load configuration
	cfg, err := config.GetConfig(".")
	if err != nil {
		if errors.Is(err, config.ErrSettingsNotFound) {
			log.Println("settings file not found, a default settings is generated and being used...")
		} else {
			log.Fatal(err)
		}
	}

	// load observer configuration
	ocfg := observerconfig.LoadObserverConfig(cfg)

	// enable window driver
	if cfg.General.Show {
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
		servoX := engine.NewServo(servos, ocfg.ServoX)
		servoY := engine.NewServo(servos, ocfg.ServoY)
		servoXY := engine.NewServos(servoX, servoY)

		// init center
		servoXY.CenterMiddleUp()
		time.Sleep(381 * time.Millisecond)

		// calibrate servos if enabled
		if cfg.Calibration.CalibrateOnStart {
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
				target.Draw(&i, ocfg.Colors.Target.ToRGBA(), ocfg.Colors.Target.T())

				// draw non-target objects
				otherObjects := append(objects[:targetX], objects[(targetX+1):]...)
				otherObjects.Draw(&i, ocfg.Colors.Other.ToRGBA(), ocfg.Colors.Other.T())

				// reset idle and suspend the counter
				if idleStatus {
					idleStatus = false
				}

				// get a target's coordinate
				lock := target.Center()

				// aim the target if it is not in the middle rectangle
				if !lock.In(ocfg.MidRect) {
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
					if time.Now().Sub(idleTime).Seconds() >= cfg.General.IdleDuration {
						fmt.Println("Idle to long ...")
						servoXY.Center()
						centered = true
					}
				}
			}

			// show window
			if cfg.General.Show {

				// draw a mid rect
				geometry.FromRect(ocfg.MidRect).Draw(&i, ocfg.Colors.MidRect.ToRGBA(), ocfg.Colors.MidRect.T())

				window.ShowImage(i)
				window.WaitKey(1)
			}

		})
	}

	// define adaptors and devices
	connections := []gobot.Connection{raspiAdaptor}
	devices := []gobot.Device{camera}

	// adds window if window is enabled
	if cfg.General.Show {
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
