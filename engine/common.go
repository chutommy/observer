package engine

import (
	"time"
)

// Center centres the Servo to 90 degree.
func (s *Servo) Center() {
	s.set(90)
}

// Center centres a duo of Servo.
func (ss *ServoXY) Center() {
	ss.servoX.set(90)
	ss.servoY.set(90)
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
