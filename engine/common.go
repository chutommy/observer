package engine

import (
	"time"
)

// center centres the Servo.
func (s *Servo) center() {
	s.set(90)
}

// Center centres a Servos.
func (ss *Servos) Center() {
	ss.servoX.center()
	ss.servoY.center()
	time.Sleep(400 * time.Millisecond)
}

// Calibrate calibrates Servos.
func (ss *Servos) Calibrate() {
	ss.Center()

	// try minDegree
	ss.servoX.set(minDegree)
	time.Sleep(400 * time.Millisecond)
	ss.servoY.set(minDegree)
	time.Sleep(400 * time.Millisecond)

	// try maxDegree
	ss.servoX.set(maxDegree)
	time.Sleep(800 * time.Millisecond)
	ss.servoY.set(maxDegree)
	time.Sleep(800 * time.Millisecond)

	ss.Center()
}

// CenterMiddleUp set Servos to the default position.
func (ss *Servos) CenterMiddleUp() {
	ss.servoX.center()
	ss.servoY.set(135)
	time.Sleep(400 * time.Millisecond)
}
