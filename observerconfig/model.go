package observerconfig

import (
	"github.com/chutommy/observer/geometry"
)

// Servo stores observer data of the servo engine.
type Servo struct {
	Pin          int64
	Calibration  float64
	Inverted     bool
	MidPoint     int
	Toleration   float64
	PxsPerDegree float64
}

// Colors stores colors of rectangles.
type Colors struct {
	Target  *geometry.Color
	Other   *geometry.Color
	MidRect *geometry.Color
}
