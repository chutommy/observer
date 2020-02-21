package main

import (
	"flag"
	"image"
	"math"
	"time"
)

// define changeable varibles
var midRect image.Rectangle
var midPoint image.Point
var pxsPerDegree float64

func init() {
	rbn := flag.String("rbname", robotName, "name of the robot")
	csc := flag.String("cascade", cascade, "path to cascade")
	srvx := flag.String("servox", servoXpin, "pin of servo controlling X axis")
	srvy := flag.String("servoy", servoYpin, "pin of servo controlling Y axis")
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
	flag.Parse()

	robotName = *rbn
	cascade = *csc
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

	resetVar()
}

func resetVar() {

	// declare mid of aim
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

	// geting safe area
	midRect = image.Rectangle{minPoint, maxPoint}

	// get number of pixels for 1 degree
	pxsPerDegree = math.Sqrt(float64(camWidth*camWidth)+float64(camHeight*camHeight)) / angleOfViewDig
}
