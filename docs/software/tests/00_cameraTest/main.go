package main

import (
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println("VideoCapture Error.")
		log.Fatal(err)
	}
	defer func(webcam *gocv.VideoCapture) {
		err := webcam.Close()
		if err != nil {
		}
	}(webcam)

	window := gocv.NewWindow("Hello")
	defer func(window *gocv.Window) {
		_ = window.Close()
	}(window)

	img := gocv.NewMat()
	defer func(img *gocv.Mat) {
		_ = img.Close()
	}(&img)

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
