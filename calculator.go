package main

import (
	"log"
)

// TODO implement reduce period
// reduce the period by the camera's max FPS property
func reducePeriod() {
	reduced := (1000 / maxFPS) + 1
	log.Printf("Reducing period from %v to %v (according to max. FPS) ...", period, reduced)
	period = reduced
}
