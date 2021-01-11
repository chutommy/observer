package main

import (
	"time"
)

// DEFAULT SETTINGS

// ==============================
// ROBOT ========================

var cascade1 = "data/frontalface_default.xml"
var cascade2 = ""

// ==============================
// SERVOS

var servoXpin int64 = 17
var servoYpin int64 = 18

// ==============================
// CAMERA =======================

var cameraSource = 0

var camWidth = 640
var camHeight = 480

var angleOfViewHor = 62.2
var angleOfViewVer = 48.8
var maxFPS time.Duration = 60

// ==============================
// PERFORMANCE ==================

var period time.Duration = 30

// ==============================
// TARGETING ====================

var aimArea = 0.15

var idleDuration float64 = 6

// ==============================
// CALIBRATION ==================

var calibration = false

var invertX = true
var invertY = true

var calibrateX = 0.7
var calibrateY = 0.5

var tolerateX float64 = 1
var tolerateY float64 = 1

var windowed = false

// ==============================
// COLORS =======================

var targetColor = cusColor{200, 30, 30, 2}
var otherColor = cusColor{20, 100, 30, 2}
var midRectColor = cusColor{20, 20, 160, 1}
