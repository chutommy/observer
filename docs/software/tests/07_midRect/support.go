package main

import (
	"image"
)

var half = aimArea / 2
var minPoint = image.Point{
	X: int(float64(midPoint.X) - float64(camWidth)*half),
	Y: int(float64(midPoint.Y) - float64(camHeight)*half),
}
var maxPoint = image.Point{
	X: int(float64(midPoint.X) + float64(camWidth)*half),
	Y: int(float64(midPoint.Y) + float64(camHeight)*half),
}
var midRect = image.Rectangle{Min: minPoint, Max: maxPoint}

func getNearestObject(objects []image.Rectangle) image.Rectangle {
	switch l := len(objects); {
	case l == 0:
		return image.Rectangle{}
	case l == 1:
		return objects[0]
	default:
		return nearestObject(objects)
	}
}

func nearestObject(rects []image.Rectangle) image.Rectangle {
	nearest := 0
	maxArea := 0

	for i, rect := range rects {
		area := (rect.Max.X - rect.Min.X) * (rect.Max.Y - rect.Min.Y)
		if area > maxArea {
			nearest = i
			maxArea = area
		}
	}

	return rects[nearest]
}

func getCoordinates(rect image.Rectangle) image.Point {
	midX := (rect.Min.X + rect.Max.X) / 2
	midY := (rect.Min.Y + rect.Max.Y) / 2
	return image.Point{
		X: midX,
		Y: midY,
	}
}
