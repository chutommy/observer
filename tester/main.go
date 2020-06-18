package test

import (
	"flag"
	"fmt"
	"time"

	piblaster "github.com/ddrager/go-pi-blaster"
)

var servos = piblaster.Blaster{}

const pin1 = 17
const pin2 = 18

var center, calibrate, test bool

// Test tests of the movement is OK
func Test() {
	servos.Start([]int64{pin1, pin2})

	time.Sleep(600 * time.Millisecond)

	if center {
		fmt.Println("Centering servos...")
		moveServo(90, 390)
	}

	if calibrate {
		moveServo(90, 390)

		moveServo(0, 390)

		moveServo(180, 780)

		moveServo(90, 390)
	}

	if test {
		moveServo(90, 390)
		moveServo(70, 250)
		moveServo(110, 300)
		moveServo(50, 350)
		moveServo(130, 400)
		moveServo(30, 450)
		moveServo(150, 500)
		moveServo(10, 550)
		moveServo(170, 600)
		moveServo(0, 650)
		moveServo(180, 700)
		moveServo(90, 400)
	}

	fmt.Println("ABOUT TO EXIT ...")
}

func init() {
	fmt.Printf("TESTING SERVOS with GPIO pins: %v, %v\n", pin1, pin2)
	cnt := (flag.Bool("center", false, "center both servos"))
	clb := (flag.Bool("calibrate", false, "calibrate both servos"))
	tst := (flag.Bool("test", false, "test the servos"))

	flag.Parse()

	center = *cnt
	calibrate = *clb
	test = *tst
}

func moveServo(angle, mils int) {
	a := float64(angle)/900 + 0.05
	servos.Apply(pin1, a)
	time.Sleep(time.Duration(mils) * time.Millisecond)
	servos.Apply(pin2, a)
	time.Sleep(time.Duration(mils) * time.Millisecond)
}
