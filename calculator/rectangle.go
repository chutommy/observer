package calculator

import (
	"image"
)

// Object is anything that is detected by the observer.
type Object struct {
	rect image.Rectangle
}

// Objects represents slice of Objects.
type Objects []*Object

// FromRect constructs an Object from a rectangle.
func FromRect(rect image.Rectangle) *Object {
	return &Object{
		rect: rect,
	}
}

// FromRects constructs a slice of Objects from the given slice of rectangles.
func FromRects(rects []image.Rectangle) Objects {
	objects := make(Objects, len(rects))

	for i, rect := range rects {
		objects[i] = &Object{
			rect: rect,
		}
	}

	return objects
}

// area calculates the area of the Object.
func (o *Object) area() int {
	return o.rect.Dx() * o.rect.Dy()
}

// Center calculates a center point of the Object.
func (o *Object) Center() image.Point {
	min := o.rect.Min
	max := o.rect.Max

	// calculate the mid point of the axis
	midX := (min.X + max.X) / 2
	midY := (min.Y + max.Y) / 2

	return image.Point{
		X: midX,
		Y: midY,
	}
}

// NearestObject returns an index of the nearest Object in the selection.
func NearestObject(objects Objects) int {
	switch len(objects) {
	case 0:
		return -1 // no objects
	case 1:
		return 0 // one objects
	default:
		return greatArea(objects) // multiple objects
	}
}

// greatArea returns an index of the Object with the greatest area value.
func greatArea(objects Objects) int {
	var maxIdx, maxArea int

	for i, object := range objects {
		a := object.area()

		if a > maxArea {
			// greater area
			maxIdx = i
			maxArea = a
		}
	}

	return maxIdx
}
