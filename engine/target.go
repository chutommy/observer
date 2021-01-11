package engine

import (
	"image"
)

// Aim moves the Servos to make sure that the given point lies
// in the tolerable rectangle area.
func (ss *Servos) Aim(point image.Point) {
	sx := ss.servoX
	sy := ss.servoY

	// calculate px differences
	xDiff := float64(point.X - sx.midPoint)
	yDiff := float64(point.Y - sy.midPoint)

	// aim on X axis
	if (xDiff > sx.toleration) || (xDiff < -sx.toleration) {
		aX := xDiff / sx.pxsPerDegree
		sx.move(aX)
	}

	// aim on Y axis
	if (yDiff > sy.toleration) || (yDiff < -sy.toleration) {
		aY := xDiff / sy.pxsPerDegree
		sy.move(aY)
	}
}
