package main

import (
	"time"
)

// ==============================
// ROBOT ========================

var robotName string = "Observing Robot"
var cascade string = "haarcascade_frontalface_default.xml"

// ==============================
// SERVOS

var servoXpin string = "1"
var servoYpin string = "2"

// ==============================
// CAMERA =======================

var cameraSource int = 0

var camWidth int = 640
var camHeight int = 480

var angleOfViewDig float64 = 69.1
var maxFPS time.Duration = 30

// ==============================
// PERFORMANCE ==================

var period time.Duration = 100

// ==============================
// TARGETING ====================

var aimArea float64 = 0.18

var idleDuration float64 = 8

// ==============================
// CALIBRATION ==================

var calibration bool = true

var invertX bool = false
var invertY bool = false

var calibrateX float64 = 1
var calibrateY float64 = 1

// ==============================
// COLORS =======================

var targetColor = cusColor{200, 30, 30, 2}
var otherColor = cusColor{20, 100, 30, 2}
var midRectColor = cusColor{20, 20, 160, 1}
