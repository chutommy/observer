package main

import (
	"image"
	"log"
)

// getNearestObject return the reactangle with the highest area
func getNearestObject(objects []image.Rectangle) (int, image.Rectangle) {
	switch l := len(objects); l {
	case 0:
		return -1, image.Rectangle{}
	case 1:
		return 0, objects[0]
	default:
		return nearestObject(objects)
	}
}

func nearestObject(rects []image.Rectangle) (int, image.Rectangle) {
	nearest := 0
	maxArea := 0

	for i, rect := range rects {
		area := (rect.Max.X - rect.Min.X) * (rect.Max.Y - rect.Min.Y)
		if area > maxArea {
			nearest = i
			maxArea = area
		}
	}

	return nearest, rects[nearest]
}

// getCoordinates returns the mid of the rectangle
func getCoordinates(rect image.Rectangle) image.Point {
	midX := (rect.Min.X + rect.Max.X) / 2
	midY := (rect.Min.Y + rect.Max.Y) / 2
	return image.Point{
		X: midX,
		Y: midY,
	}
}

// reducePeriod reduces periods time if it is too low
func reducePeriod() {
	reduced := 1000 / maxFPS
	log.Printf("Reducing period from %v to %v (max. FPS) ...", period, reduced)
	period = reduced
}
