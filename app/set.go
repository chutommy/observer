package main

import (
	"time"
)

// DEFAULT SETTINGS

// ==============================
// ROBOT ========================

var robotName string = "Observing Robot"

var cascade1 string = "data/frontalface_default.xml"
var cascade2 string = ""

// ==============================
// SERVOS

var servoXpin int64 = 17
var servoYpin int64 = 18

// ==============================
// CAMERA =======================

var cameraSource int = 0

var camWidth int = 640
var camHeight int = 480

var angleOfViewHor float64 = 62.2
var angleOfViewVer float64 = 48.8
var maxFPS time.Duration = 60

// ==============================
// PERFORMANCE ==================

var period time.Duration = 30

// ==============================
// TARGETING ====================

var aimArea float64 = 0.15

var idleDuration float64 = 6

// ==============================
// CALIBRATION ==================

var calibration bool = true

var invertX bool = true
var invertY bool = true

var calibrateX float64 = 1
var calibrateY float64 = 1

// ==============================
// COLORS =======================

var windowed bool = false

// ==============================
// COLORS =======================

var targetColor = cusColor{200, 30, 30, 2}
var otherColor = cusColor{20, 100, 30, 2}
var midRectColor = cusColor{20, 20, 160, 1}
