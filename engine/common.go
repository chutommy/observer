package engine

import (
	"time"
)

// center centres the Servo to 90 degree.
func (s *Servo) center() {
	s.set(90)
}

// Center centres a Servos.
func (ss *Servos) Center() {
	ss.servoX.center()
	ss.servoY.center()
}

// Calibrate calibrates Servos.
func (ss *Servos) Calibrate() {
	ss.Center()
	time.Sleep(400 * time.Millisecond)

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
	time.Sleep(400 * time.Millisecond)
}

// CenterMiddleUp set Servos to the default position.
func (ss *Servos) CenterMiddleUp() {
	ss.servoX.center()
	ss.servoY.set(135)
}
