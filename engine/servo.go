package engine

import (
	piblaster "github.com/ddrager/go-pi-blaster"
)

type Servo struct {
	blaster      piblaster.Blaster
	degreeStatus int
}
