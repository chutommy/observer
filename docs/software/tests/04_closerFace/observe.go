package main

import "image"

func noDetect() {}

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
