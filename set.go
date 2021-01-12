package main

import (
	"time"
)

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
