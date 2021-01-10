package main

import (
	"flag"
	"image"
	"time"

	piblaster "github.com/ddrager/go-pi-blaster"
)

// define variables that might change and make them global
var midRect image.Rectangle
var midPoint image.Point
var pxsPerDegreeVer float64
var pxsPerDegreeHor float64
var servos = piblaster.Blaster{}
var tolerationX float64
var tolerationY float64
var tolerationXr float64
var tolerationYr float64

func init() {
	csc1 := flag.String("casc1", cascade1, "path to cascade")
	csc2 := flag.String("casc2", cascade2, "path to cascade (optional)")
	srvx := flag.Int64("servox", servoXpin, "GPIO pin of servoX (not the pin number)")
	srvy := flag.Int64("servoy", servoYpin, "GPIO pin of servoY (not the pin number)")
	cams := flag.Int("camsource", cameraSource, "camera source")
	camw := flag.Int("camwidth", camWidth, "camera width")
	camh := flag.Int("camheight", camHeight, "camera height")
	aovhor := flag.Float64("angleovh", angleOfViewHor, "camera's horizontal angle of view")
	aovver := flag.Float64("angleovv", angleOfViewVer, "camera's vertical angle of view")
	mfps := flag.Int64("maxfps", int64(maxFPS), "camera's max FPS")
	prd := flag.Int64("period", int64(period), "speed of the loop (ns)")
	aima := flag.Float64("aimarea", aimArea, "aim area (1~100%)")
	idldur := flag.Float64("idledur", idleDuration, "time until reset (sec)")
	calib := flag.Bool("calib", calibration, "calibration on start")
	invx := flag.Bool("invertx", invertX, "invert X")
	invy := flag.Bool("inverty", invertY, "invert Y")
	calibx := flag.Float64("calibx", calibrateX, "calibrate X (1~100%)")
	caliby := flag.Float64("caliby", calibrateY, "calibrate Y (1~100%)")
	tx := flag.Float64("tolx", tolerateX, "toleration of X (1~100%)")
	ty := flag.Float64("toly", tolerateY, "toleration of Y (1~100%)")
	// window enable
	wnd := flag.Bool("window", windowed, "enable window (ONLY IF DISPLAY OUTPUT IS AVAILABLE) - ensure VNC or HDMI output")
	flag.Parse()

	cascade1 = *csc1
	cascade2 = *csc2
	servoXpin = *srvx
	servoYpin = *srvy
	cameraSource = *cams
	camWidth = *camw
	camHeight = *camh
	angleOfViewHor = *aovhor
	angleOfViewVer = *aovver
	maxFPS = time.Duration(*mfps)
	period = time.Duration(*prd)
	aimArea = *aima
	idleDuration = *idldur
	calibration = *calib
	invertX = *invx
	invertY = *invy
	calibrateX = *calibx
	calibrateY = *caliby
	tolerateX = *tx
	tolerateY = *ty
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

	// get toleration
	tolerationX = (float64(midRect.Dx()) / 2) * tolerateX
	tolerationY = (float64(midRect.Dy()) / 2) * tolerateY
	tolerationXr -= tolerationX
	tolerationYr -= tolerationY

	// get number of pixels for 1 degree
	pxsPerDegreeHor = float64(camWidth) / angleOfViewHor
	pxsPerDegreeVer = float64(camHeight) / angleOfViewVer

	// initiate servo drivers
	servos.Start([]int64{servoXpin, servoYpin})
}