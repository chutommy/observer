package engine

import (
	"math"

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

	calibration  float64
	midPoint     int
	toleration   float64
	pxsPerDegree float64
}

// NewServo is a constructor of the Servo.
func NewServo(blaster blaster.Blaster, pin int64, inverted bool, calibration float64, midPoint int, toleration float64, pxsPerDegree float64) *Servo {
	// set calibration for inversion
	if inverted {
		calibration = -calibration
	}

	// construct
	s := &Servo{
		blaster:      blaster,
		degreeStatus: 90,
		pin:          pin,
		calibration:  calibration,
		midPoint:     midPoint,
		toleration:   toleration,
		pxsPerDegree: pxsPerDegree,
	}

	// center
	s.center()

	return s
}

// Servos represent a duo of Servo.
type Servos struct {
	servoX *Servo
	servoY *Servo
}

// NewServos is a constructor of the Servos.
func NewServos(sX, sY *Servo) *Servos {
	// construct
	ss := &Servos{
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

// move changes the angle of the Servo.
// Respect the maximum/minimum range.
func (s *Servo) move(angle float64) {
	// calibration (+ inversion)
	angle *= s.calibration

	// movement range
	newAngle := s.degreeStatus + int(math.Round(angle))

	switch {
	case newAngle < minDegree:
		s.set(minDegree)
	case newAngle > maxDegree:
		s.set(maxDegree)
	default:
		s.set(newAngle)
	}
}
