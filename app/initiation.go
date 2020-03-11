package main

import (
	"flag"
	"image"
	"math"
	"time"

	piblaster "github.com/ddrager/go-pi-blaster"
)

// define variables that might change and make them global
var midRect image.Rectangle
var midPoint image.Point
var pxsPerDegree float64
var servos = piblaster.Blaster{}

func init() {
	rbn := flag.String("rbname", robotName, "name of the robot")
	csc1 := flag.String("cascade 1", cascade1, "path to cascade")
	csc2 := flag.String("cascade 2", cascade2, "path to cascade (optional)")
	srvx := flag.Int64("servox", servoXpin, "GPIO pin of servo controlling X axis (not the pin number)")
	srvy := flag.Int64("servoy", servoYpin, "GPIO pin of servo controlling Y axis (not the pin number)")
	cams := flag.Int("camsource", cameraSource, "camera source")
	camw := flag.Int("camwidth", camWidth, "camera width")
	camh := flag.Int("camheight", camHeight, "camera height")
	aov := flag.Float64("angleov", angleOfViewDig, "camera's diagonal angle of view")
	mfps := flag.Int64("maxfps", int64(maxFPS), "camera's maximal FPS")
	prd := flag.Int64("period", int64(period), "speed of shooting in ns")
	aima := flag.Float64("aimarea", aimArea, "aim area (0-0%, 1-100%)")
	idldur := flag.Float64("idledur", idleDuration, "duration of not detecting faces")
	calib := flag.Bool("calib", calibration, "calibration on start")
	invx := flag.Bool("invertx", invertX, "invert X aiming")
	invy := flag.Bool("inverty", invertY, "invert Y aiming")
	calibx := flag.Float64("calibx", calibrateX, "calibrate X")
	caliby := flag.Float64("caliby", calibrateY, "calibrate Y")
	// window enabling
	wnd := flag.Bool("window", windowed, "enable window (ONLY IF DISPLAY IS AVAILABLE) - ensure VNC or HDMI output")
	flag.Parse()

	robotName = *rbn
	cascade1 = *csc1
	cascade2 = *csc2
	servoXpin = *srvx
	servoYpin = *srvy
	cameraSource = *cams
	camWidth = *camw
	camHeight = *camh
	angleOfViewDig = *aov
	maxFPS = time.Duration(*mfps)
	period = time.Duration(*prd)
	aimArea = *aima
	idleDuration = *idldur
	calibration = *calib
	invertX = *invx
	invertY = *invy
	calibrateX = *calibx
	calibrateY = *caliby
	windowed = *wnd

	resetVar()
}

func resetVar() {

	// declare the center of the aiming screen
	midPoint = image.Point{
		X: camWidth / 2,
		Y: camHeight / 2,
	}

	// get an aiming area
	half := aimArea / 2
	minPoint := image.Point{
		int(float64(midPoint.X) - float64(camWidth)*half),
		int(float64(midPoint.Y) - float64(camWidth)*half),
	}
	maxPoint := image.Point{
		int(float64(midPoint.X) + float64(camWidth)*half),
		int(float64(midPoint.Y) + float64(camWidth)*half),
	}
	midRect = image.Rectangle{minPoint, maxPoint}

	// get number of pixels for 1 degree
	pxsPerDegree = math.Sqrt(float64(camWidth*camWidth)+float64(camHeight*camHeight)) / angleOfViewDig

	// initiate servo drivers
	servos.Start([]int64{servoXpin, servoYpin})
}
