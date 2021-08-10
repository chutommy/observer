package main

import (
	"image"
	"log"

	"gobot.io/x/gobot/platforms/opencv"
	"gocv.io/x/gocv"
)

var midPoint = image.Point{
	X: camWidth / 2,
	Y: camHeight / 2,
}

var midRect image.Rectangle

var half = aimArea / 2
var minPoint = image.Point{
	X: int(float64(midPoint.X) - float64(camWidth)*half),
	Y: int(float64(midPoint.Y) - float64(camWidth)*half),
}
var maxPoint = image.Point{
	X: int(float64(midPoint.X) + float64(camWidth)*half),
	Y: int(float64(midPoint.Y) + float64(camWidth)*half),
}

func init() {
	midRect = image.Rectangle{Min: minPoint, Max: maxPoint}
}

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

func getCoordinates(rect image.Rectangle) image.Point {
	midX := (rect.Min.X + rect.Max.X) / 2
	midY := (rect.Min.Y + rect.Max.Y) / 2
	return image.Point{
		X: midX,
		Y: midY,
	}
}

func sumCascades(img gocv.Mat) []image.Rectangle {
	var sum []image.Rectangle

	for _, cascade := range cascades {
		sum = append(sum, opencv.DetectObjects(cascade, img)...)
	}

	return sum
}

func reducePeriod() {
	reduced := 1000 / maxFPS
	log.Printf("Reducing period from %v to %v (max. FPS) ...", period, reduced)
	period = reduced
}
