package main

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func startObserve(camID int, xmlFile string) {

	cam, err := gocv.VideoCaptureDevice(camID)
	if err != nil {
		fmt.Printf("Error: cannot load cam, device ID: %v\n", camID)
		return
	}
	defer func(cam *gocv.VideoCapture) {
		_ = cam.Close()
	}(cam)

	img := gocv.NewMat()
	defer func(img *gocv.Mat) {
		_ = img.Close()
	}(&img)

	classifier := gocv.NewCascadeClassifier()
	defer func(classifier *gocv.CascadeClassifier) {
		_ = classifier.Close()
	}(&classifier)
	if !classifier.Load(xmlFile) {
		fmt.Printf("Error: cannot load cascade (xml) file, filename: %v\n", xmlFile)
		return
	}

	observeLoop(cam, img, classifier)
}

func observeLoop(cam *gocv.VideoCapture, img gocv.Mat, classifier gocv.CascadeClassifier) {
	fmt.Printf("start reading camera device\n")

	for {
		if ok := cam.Read(&img); !ok {
			fmt.Printf("Error: cannot read cam device\n")
			return
		}

		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		var target image.Rectangle
		if l := len(rects); l == 1 {
			target = rects[0]
		} else if l > 1 {
			target = nearestObject(rects)
		} else if l < 1 {
			noDetect()
			continue
		}

		midX := (target.Min.X + target.Max.X) / 2
		midY := (target.Min.Y + target.Max.Y) / 2
		lock := image.Point{
			X: midX,
			Y: midY,
		}

		doAction(lock)
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

func noDetect() {}

func doAction(pt image.Point) {
	fmt.Printf("X: %v\tY: %v\n", pt.X, pt.Y)
}
