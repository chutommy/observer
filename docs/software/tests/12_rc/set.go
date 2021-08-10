package main

import (
	"time"
)

var robotName = "Observing Robot"
var cascades = []string{
	"haarcascade_frontalface_default.xml",
}

var servoXpin = "1"
var servoYpin = "2"

var cameraSource = 0

var camWidth = 640
var camHeight = 480

var angleOfViewDig = 69.1
var maxFPS time.Duration = 30

var period time.Duration = 100

var aimArea = 0.18

var idleDuration float64 = 8

var targetColor = cusColor{200, 30, 30, 2}
var otherColor = cusColor{20, 100, 30, 2}
var midRectColor = cusColor{20, 20, 160, 1}

var calibration = true

var invertX float64 = 1
var invertY float64 = 1

var calibrateX float64 = 0
var calibrateY float64 = 0
