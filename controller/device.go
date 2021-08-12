package controller

import (
	"gobot.io/x/gobot"
)

// LoadRobot loads all required devices for the Robot.
func (o *Observer) LoadRobot() {
	o.log.Info("Initiating adaptors and connections")

	// adaptors and devices
	conns := gobot.Connections{o.adaptor}
	devices := gobot.Devices{o.camera}

	if o.cfg.Show {
		devices = append(devices, o.window)
	}

	o.robot = gobot.NewRobot(
		o.name,
		conns,
		devices,
		o.work,
	)
}

// Start starts the device.
func (o *Observer) Start() error {
	return o.robot.Start()
}
