package main

import (
	"log"
)

// TODO implement others

// color type.
type cusColor struct {
	r, g, b, thickness int
}

// reduce the period by the camera's max FPS property.
func reducePeriod() {
	reduced := (1000 / maxFPS) + 1
	log.Printf("Reducing period from %v to %v (according to max. FPS) ...", period, reduced)
	period = reduced
}
