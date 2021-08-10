package main

import (
	"fmt"
	"image"
	"image/color"
	"runtime"

	"gocv.io/x/gocv"
)

func main() {
	runtime.GOMAXPROCS(1)

	deviceID := 0
	xmlFile := "../haarcascades/haarcascade_frontalface_default.xml"

	webcam, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		fmt.Printf("Webcam load error. Device ID: %v\n", deviceID)
		return
	}
	defer func(webcam *gocv.VideoCapture) {
		_ = webcam.Close()
	}(webcam)

	window := gocv.NewWindow("Face Detect")
	defer func(window *gocv.Window) {
		_ = window.Close()
	}(window)

	img := gocv.NewMat()
	defer func(img *gocv.Mat) {
		_ = img.Close()
	}(&img)

	red := color.RGBA{R: 255}
	black := color.RGBA{}

	classifier := gocv.NewCascadeClassifier()
	defer func(classifier *gocv.CascadeClassifier) {
		_ = classifier.Close()
	}(&classifier)

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		for _, r := range rects {
			gocv.Rectangle(&img, r, red, 3)
			gocv.Rectangle(&img, r, red, 3)

			size := gocv.GetTextSize("Opice", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, "Opice", pt, gocv.FontHersheyPlain, 1.2, black, 2)
		}

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
