package engine

import (
	"math"

	"github.com/chutommy/observer/observerconfig"

	blaster "github.com/ddrager/go-pi-blaster"
)

const (
	minDegree = 0
	maxDegree = 180
)

// Servo represents a Servo engine controller.
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
func NewServo(blaster blaster.Blaster, servo *observerconfig.Servo) *Servo {
	if servo.Inverted {
		servo.Calibration *= -1
	}

	s := &Servo{
		blaster:      blaster,
		degreeStatus: 90,
		pin:          servo.Pin,
		calibration:  servo.Calibration,
		midPoint:     servo.MidPoint,
		toleration:   servo.Toleration,
		pxsPerDegree: servo.PxsPerDegree,
	}
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
	ss := &Servos{
		servoX: sX,
		servoY: sY,
	}
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

// move changes the angle of the Servo. It respects the maximum/minimum range.
func (s *Servo) move(angle float64) {
	angle *= s.calibration
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
