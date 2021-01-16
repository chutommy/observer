package controller

import (
	"gobot.io/x/gobot"
)

// LoadRobot loads all required devices for the Robot.
func (o *Observer) LoadRobot() {
	// define adaptors and devices
	conns := gobot.Connections{o.adaptor}
	devices := gobot.Devices{o.camera}

	// add window if Show is enabled
	if o.cfg.Show {
		devices = append(devices, o.window)
	}

	// init robot
	o.robot = gobot.NewRobot(
		o.name,
		conns,
		devices,
		o.work,
	)
}

// Start starts the Observer.
func (o *Observer) Start() error {
	return o.robot.Start()
}
