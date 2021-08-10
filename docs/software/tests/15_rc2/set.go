package main

import (
	"time"
)

var robotName = "Observing Robot"
var cascade = "haarcascade_frontalface_default.xml"

var servoXpin int64 = 17
var servoYpin int64 = 18

var cameraSource = 0

var camWidth = 640
var camHeight = 480

var angleOfViewDig = 79.05
var maxFPS time.Duration = 30

var period time.Duration = 45

var aimArea = 0.18

var idleDuration float64 = 12

var calibration = false

var invertX = false
var invertY = false

var calibrateX float64 = 1
var calibrateY float64 = 1

var windowed = false

var targetColor = cusColor{200, 30, 30, 2}
var otherColor = cusColor{20, 100, 30, 2}
var midRectColor = cusColor{20, 20, 160, 1}
