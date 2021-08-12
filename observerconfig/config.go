package observerconfig

import (
	"image"

	"observer/config"
	"observer/geometry"
)

// ObserverConfig represents a configuration of the Observer's runtime.
type ObserverConfig struct {
	MidPoint image.Point
	MidRect  image.Rectangle
	ServoX   *Servo
	ServoY   *Servo
	Colors   *Colors

	Show             bool
	CalibrateOnStart bool
	Haar             []string
	MaxIdleDuration  float64
	MaxFPS           int
	Period           int
}

// LoadObserverConfig generates an ObserverConfig and loads required attributes.
func LoadObserverConfig(cfg *config.Config) *ObserverConfig {
	rc := &ObserverConfig{}

	rc.loadMidPoint(cfg)
	rc.loadMidRect(cfg)
	rc.loadServoX(cfg)
	rc.loadServoY(cfg)
	rc.loadColors(cfg)
	rc.loadGeneral(cfg)

	return rc
}

// loadMidPoint sets MidPoint.
func (rc *ObserverConfig) loadMidPoint(cfg *config.Config) {
	frame := cfg.Camera.Frame
	rc.MidPoint = image.Point{
		X: frame.Width / 2,
		Y: frame.Height / 2,
	}
}

// loadMidRect sets MidRect.
func (rc *ObserverConfig) loadMidRect(cfg *config.Config) {
	half := cfg.Targeting.AimArea / 2
	wh := float64(cfg.Camera.Frame.Width) * half
	minPoint := image.Point{
		X: int(float64(rc.MidPoint.X) - wh),
		Y: int(float64(rc.MidPoint.Y) - wh),
	}
	maxPoint := image.Point{
		X: int(float64(rc.MidPoint.X) + wh),
		Y: int(float64(rc.MidPoint.Y) + wh),
	}

	rc.MidRect = image.Rectangle{
		Min: minPoint,
		Max: maxPoint,
	}
}

// loadServoX sets ServoX.
func (rc *ObserverConfig) loadServoX(cfg *config.Config) {
	rc.ServoX = &Servo{
		Pin:          int64(cfg.Servos.PinX),
		Calibration:  cfg.Calibration.Coefficient.X,
		Inverted:     cfg.Calibration.Invert.X,
		MidPoint:     cfg.Camera.Frame.Width / 2,
		Toleration:   (float64(rc.MidRect.Dx()) / 2) * cfg.Calibration.Tolerate.X,
		PxsPerDegree: float64(cfg.Camera.Frame.Width) / cfg.Camera.AngleOfView.Horizontal,
	}
}

// loadServoY sets ServoY.
func (rc *ObserverConfig) loadServoY(cfg *config.Config) {
	rc.ServoY = &Servo{
		Pin:          int64(cfg.Servos.PinY),
		Calibration:  cfg.Calibration.Coefficient.Y,
		Inverted:     cfg.Calibration.Invert.Y,
		MidPoint:     cfg.Camera.Frame.Height / 2,
		Toleration:   (float64(rc.MidRect.Dy()) / 2) * cfg.Calibration.Tolerate.Y,
		PxsPerDegree: float64(cfg.Camera.Frame.Height) / cfg.Camera.AngleOfView.Vertical,
	}
}

// loadColors sets Colors.
func (rc *ObserverConfig) loadColors(cfg *config.Config) {
	color := cfg.Targeting.Color
	t, o, m := color.Target, color.Other, color.MidRect

	rc.Colors = &Colors{
		Target:  geometry.NewColor(t.Red, t.Green, t.Blue, t.Thickness),
		Other:   geometry.NewColor(o.Red, o.Green, o.Blue, o.Thickness),
		MidRect: geometry.NewColor(m.Red, m.Green, m.Blue, m.Thickness),
	}
}

// loadGeneral sets general variables.
func (rc *ObserverConfig) loadGeneral(cfg *config.Config) {
	rc.Show = cfg.General.Show
	rc.CalibrateOnStart = cfg.Calibration.CalibrateOnStart
	rc.Haar = cfg.Targeting.Cascades
	rc.MaxIdleDuration = cfg.General.IdleDuration
	rc.MaxFPS = cfg.Camera.MaxFPS
	rc.Period = cfg.General.Period
}
