package main

import (
	piblaster "github.com/ddrager/go-pi-blaster"
)

// define variables that might change and make them global
var servos = piblaster.Blaster{}

func init() {
	resetVar()
}

func resetVar() {

	// initiate servo drivers
	servos.Start([]int64{servoXpin, servoYpin})
}
