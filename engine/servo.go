package engine

import (
	blaster "github.com/ddrager/go-pi-blaster"
)

const (
	minDegree = 0
	maxDegree = 180
)

// Servo represents a single Servo engine controller.
type Servo struct {
	blaster      blaster.Blaster
	degreeStatus int
	pin          int64
}

// NewServo is a constructor of the Servo.
func NewServo(b blaster.Blaster, pin int64) *Servo {
	// construct
	s := &Servo{
		blaster:      b,
		degreeStatus: 90,
		pin:          pin,
	}

	// center
	s.Center()

	return s
}

// ServoXY represent a duo of Servo.
type ServoXY struct {
	servoX *Servo
	servoY *Servo
}

// NewServoXY is a constructor of the ServoXY.
func NewServoXY(sX, sY *Servo) *ServoXY {
	// construct
	ss := &ServoXY{
		servoX: sX,
		servoY: sY,
	}

	// center
	ss.Center()

	return ss
}

// set sets the Servo to the specific angle.
func (s *Servo) set(angle int) {
	s.degreeStatus = angle

	// PWD calculation
	a := float64(angle)/900 + 0.05

	// change/apply PWD signal
	s.blaster.Apply(s.pin, a)
}
