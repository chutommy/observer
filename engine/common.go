package engine

import (
	"time"
)

// center centres the servo to 90 degree.
func (s *servo) center() {
	s.set(90)
}

// Center centres a servos.
func (ss *servos) Center() {
	ss.servoX.center()
	ss.servoY.center()
}

// Calibrate tries the range for the Servo.
func (s *Servo) Calibrate() {
	s.set(minDegree)
	time.Sleep(800 * time.Millisecond)
	s.set(maxDegree)
	time.Sleep(800 * time.Millisecond)
}

// Calibrate calibrates both Servo.
func (ss *ServoXY) Calibrate() {
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

// CenterMiddleUp set servos to the default position.
func (ss *servos) CenterMiddleUp() {
	ss.servoX.center()
	ss.servoY.set(135)
}
