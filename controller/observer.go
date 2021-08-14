package controller

import (
	"sync/atomic"
	"time"

	"github.com/chutommy/observer/config"
	"github.com/chutommy/observer/engine"
	"github.com/chutommy/observer/observerconfig"

	blaster "github.com/ddrager/go-pi-blaster"
	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/opencv"
	"gobot.io/x/gobot/platforms/raspi"
	"gocv.io/x/gocv"
)

// Observer represents the robot's controller.
type Observer struct {
	robot *gobot.Robot

	name string
	log  *logrus.Entry
	cfg  *observerconfig.ObserverConfig

	// devices
	adaptor *raspi.Adaptor
	camera  *opencv.CameraDriver
	window  *opencv.WindowDriver

	// servos
	blaster blaster.Blaster
	servos  *engine.Servos

	work         func()
	activeFrame  *atomic.Value
	currentFrame *gocv.Mat

	lastUpdated time.Time
	idle        bool
}

// NewObserver constructs a new Observer controller.
func NewObserver(name string, log *logrus.Entry, cfg *config.Config) *Observer {
	o := &Observer{
		robot: nil,

		name: name,
		log:  log,
		cfg:  observerconfig.LoadObserverConfig(cfg),

		adaptor: raspi.NewAdaptor(),
		camera:  opencv.NewCameraDriver(cfg.Camera.Source),
		window:  nil, // optional

		blaster: blaster.Blaster{},
		servos:  nil,

		work:         func() {},
		activeFrame:  &atomic.Value{}, // live frame
		currentFrame: &gocv.Mat{},     // last loaded frame
	}

	o.initWindow()
	o.initServos()
	o.initCamera()
	o.checkFrequency()

	o.log.Info("Observer internal structure successfully constructed")

	return o
}
