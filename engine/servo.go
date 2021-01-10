package engine

import (
	"math"

	blaster "github.com/ddrager/go-pi-blaster"
)

const (
	minDegree = 0
	maxDegree = 180
)

// servo represents a single servo engine controller.
type servo struct {
	blaster      blaster.Blaster
	degreeStatus int
	pin          int64

	calibration  float64
	midPoint     int
	toleration   float64
	pxsPerDegree float64
}

// NewServo is a constructor of the servo.
func NewServo(blaster blaster.Blaster, pin int64, inverted bool, calibration float64, midPoint int, toleration float64, pxsPerDegree float64) *servo {
	// set calibration for inversion
	if inverted {
		calibration = -calibration
	}

	// construct
	s := &servo{
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

// servos represent a duo of servo.
type servos struct {
	servoX *servo
	servoY *servo
}

// NewServos is a constructor of the servos.
func NewServos(sX, sY *servo) *servos {
	// construct
	ss := &servos{
		servoX: sX,
		servoY: sY,
	}

	// center
	ss.Center()

	return ss
}

// set sets the servo to the specific angle.
func (s *servo) set(angle int) {
	s.degreeStatus = angle

	// PWD calculation
	a := float64(angle)/900 + 0.05

	// change/apply PWD signal
	s.blaster.Apply(s.pin, a)
}
