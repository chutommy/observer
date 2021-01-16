package controller

import (
	"observer/engine"

	"gobot.io/x/gobot/platforms/opencv"
	"gocv.io/x/gocv"
)

// initWindow sets window driver if Show option is set up.
func (o *Observer) initWindow() {
	if o.cfg.Show {
		o.window = opencv.NewWindowDriver()
	}
}

// initServos initializes the blaster and connected servo motors.
func (o *Observer) initServos() {
	o.blaster.Start([]int64{o.cfg.ServoX.Pin, o.cfg.ServoY.Pin})

	servoX := engine.NewServo(o.blaster, o.cfg.ServoX)
	servoY := engine.NewServo(o.blaster, o.cfg.ServoY)
	o.servos = engine.NewServos(servoX, servoY)
}

// initCamera turns on the camera and initializes the connection
// between the record and current active frame.
func (o *Observer) initCamera() {
	// prepare a structure for an active frame
	mat := gocv.NewMat()
	o.activeFrame.Store(mat)

	// turn camera on
	_ = o.camera.On(opencv.Frame, func(data interface{}) {
		cam := data.(gocv.Mat)
		o.activeFrame.Store(cam)
	})
}

// checkFrequency validates the value of the period and if the value is
// unnecessarily high and unmaintainable, it is automatically reduced.
func (o *Observer) checkFrequency() {
	if 1000/o.cfg.Period > o.cfg.MaxFPS {
		o.cfg.Period = (1000 / o.cfg.MaxFPS) + 1
	}
}
