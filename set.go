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
var maxFPS time.Duration = 60

// ==============================
// PERFORMANCE ==================

var period time.Duration = 30
