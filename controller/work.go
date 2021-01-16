package controller

import (
	"time"

	"observer/geometry"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
	"gocv.io/x/gocv"
)

// LoadWork sets a work attribute for the Observer.
func (o *Observer) LoadWork() {
	o.work = func() {
		o.log.Info("Starting the Observer")

		// calibrate if set
		if o.cfg.CalibrateOnStart {
			o.log.Info("Calibrating before the observing cycle initiation")

			o.servos.Center()
			o.servos.Calibrate()
		}

		// center
		o.servos.CenterMiddleUp()

		o.log.Info("Starting the Observer's cycle")

		// start the observing cycle
		gobot.Every(time.Duration(o.cfg.Period)*time.Millisecond, func() {
			o.observeCycle()

			// show the frame
			if o.cfg.Show {
				geometry.FromRect(o.cfg.MidRect).Draw(o.currentFrame, o.cfg.Colors.MidRect.ToRGBA(), o.cfg.Colors.MidRect.T())
				o.window.ShowImage(*o.currentFrame)
			}
		})
	}
}

// observeCycle defines an action which is Observer repeating every period.
func (o *Observer) observeCycle() {
	// load frame
	if !o.loadFrame() {
		return
	}

	// scan objects
	objects := o.scanObjects()

	// get target
	targetX := geometry.NearestObject(objects) // obtain target's index
	if targetX == -1 {
		if time.Since(o.lastUpdated) >= time.Duration(o.cfg.MaxIdleDuration) && !o.idle {
			o.idle = true
			o.servos.CenterMiddleUp()
		}

		return
	}

	// draw target's rectangle
	target := objects[targetX]
	target.Draw(o.currentFrame, o.cfg.Colors.Target.ToRGBA(), o.cfg.Colors.Target.T())

	// draw other rectangles
	otherObjects := append(objects[:targetX], objects[targetX+1:]...)
	otherObjects.Draw(o.currentFrame, o.cfg.Colors.MidRect.ToRGBA(), o.cfg.Colors.Other.T())

	// get a target's center
	lock := target.Center()

	// aim the target if not inside the MidRect
	if !lock.In(o.cfg.MidRect) {
		o.servos.Aim(lock)
	}

	// update lastUpdated attribute
	o.lastUpdated = time.Now()
}

// loadFrame loads the current frame from the current frame
// and returns true if frame is successfully retrieved.
func (o *Observer) loadFrame() bool {
	o.currentFrame = o.activeFrame.Load().(*gocv.Mat)

	return !o.currentFrame.Empty()
}

// scanObjects scans the current frame and returns detected Objects.
func (o *Observer) scanObjects() geometry.Objects {
	objects := make(geometry.Objects, 0)

	// append scanned objects
	for _, haar := range o.cfg.Haar {
		rects := opencv.DetectObjects(haar, *o.currentFrame)
		objects = append(objects, geometry.FromRects(rects)...)
	}

	return objects
}
