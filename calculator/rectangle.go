package calculator

import (
	"image"
)

// NearestRect returns an index of the nearest rectangle in the selection.
func NearestRect(rects []image.Rectangle) int {
	switch len(rects) {
	case 0:
		return -1 // no rectangles
	case 1:
		return 0 // one rectangle
	default:
		return greatArea(rects) // multiple rectangles
	}
}

// greatArea returns an index of the rectangle with the greatest area value.
func greatArea(rects []image.Rectangle) int {
	var maxIdx, maxArea int

	for i, rect := range rects {
		a := area(rect)

		if a > maxArea {
			maxIdx = i
			maxArea = a
		}
	}

	return maxIdx
}

// area calculates the area of the rectangle.
func area(r image.Rectangle) int {
	return r.Dx() * r.Dy()
}

// Center calculates a Center of the rectangle.
func Center(r image.Rectangle) image.Point {
	min := r.Min
	max := r.Max

	// calculate mid on axis
	midX := (min.X + max.X) / 2
	midY := (min.Y + max.Y) / 2

	return image.Point{
		X: midX,
		Y: midY,
	}
}
