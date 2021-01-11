package main

import (
	"image"

	piblaster "github.com/ddrager/go-pi-blaster"
)

// define variables that might change and make them global
var midRect image.Rectangle
var midPoint image.Point
var pxsPerDegreeVer float64
var pxsPerDegreeHor float64
var servos = piblaster.Blaster{}
var tolerationX float64
var tolerationY float64
var tolerationXr float64
var tolerationYr float64

func resetVar() {

	// declare the center of the aiming screen
	midPoint = image.Point{
		X: camWidth / 2,
		Y: camHeight / 2,
	}

	// get an aiming area
	half := aimArea / 2
	minPoint := image.Point{
		int(float64(midPoint.X) - float64(camWidth)*half),
		int(float64(midPoint.Y) - float64(camWidth)*half),
	}
	maxPoint := image.Point{
		int(float64(midPoint.X) + float64(camWidth)*half),
		int(float64(midPoint.Y) + float64(camWidth)*half),
	}
	midRect = image.Rectangle{minPoint, maxPoint}

	// get toleration
	tolerationX = (float64(midRect.Dx()) / 2) * tolerateX
	tolerationY = (float64(midRect.Dy()) / 2) * tolerateY
	tolerationXr -= tolerationX
	tolerationYr -= tolerationY

	// get number of pixels for 1 degree
	pxsPerDegreeHor = float64(camWidth) / angleOfViewHor
	pxsPerDegreeVer = float64(camHeight) / angleOfViewVer

	// initiate servo drivers
	servos.Start([]int64{servoXpin, servoYpin})
}
