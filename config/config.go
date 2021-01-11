package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// ErrSettingsNotFound is returned if settings file is not found.
var ErrSettingsNotFound = errors.New("settings file not found, default configuration is being generated and used")

// GetConfig gets the configuration file for the observer.
func GetConfig(path string) (*Config, error) {
	v := viper.New()

	// set default
	setDefault(v)

	// load from file
	v.SetConfigName("settings")
	v.SetConfigType("toml")
	v.AddConfigPath(path)

	// read
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			cfg, _ := toConfig(v)

			// generate file
			_ = v.SafeWriteConfig()

			return cfg, ErrSettingsNotFound
		}

		return nil, fmt.Errorf("could not read the file: %w", err)
	}

	cfg, err := toConfig(v)
	if err != nil {
		return nil, fmt.Errorf("unable to set configuration: %w", err)
	}

	return cfg, nil
}

// setDefault sets default values of the settings for the viper Config.
func setDefault(v *viper.Viper) {
	// general
	v.SetDefault("general.show", false)
	v.SetDefault("general.period", 30)
	v.SetDefault("general.idleDuration", 6)

	// servos
	v.SetDefault("servos.pinX", 17)
	v.SetDefault("servos.pinY", 18)

	// camera
	v.SetDefault("camera.source", 0)
	v.SetDefault("camera.maxFPS", 60)
	v.SetDefault("camera.frame.width", 640)
	v.SetDefault("camera.frame.height", 480)
	v.SetDefault("camera.angleOfView.horizontal", 62.2)
	v.SetDefault("camera.angleOfView.vertical", 48.8)
	// targeting
	v.SetDefault("targeting.aimArea", 0.15)
	v.SetDefault("targeting.cascades", []string{"data/frontalface_default.xml"})

	v.SetDefault("targeting.color.target.red", 200)
	v.SetDefault("targeting.color.target.green", 30)
	v.SetDefault("targeting.color.target.blue", 30)
	v.SetDefault("targeting.color.target.thickness", 2)

	v.SetDefault("targeting.color.other.red", 20)
	v.SetDefault("targeting.color.other.green", 100)
	v.SetDefault("targeting.color.other.blue", 30)
	v.SetDefault("targeting.color.other.thickness", 2)

	v.SetDefault("targeting.color.midRect.red", 20)
	v.SetDefault("targeting.color.midRect.green", 20)
	v.SetDefault("targeting.color.midRect.blue", 160)
	v.SetDefault("targeting.color.midRect.thickness", 1)

	// calibration
	v.SetDefault("calibration.calibrateOnStart", false)
	v.SetDefault("calibration.invert.x", true)
	v.SetDefault("calibration.invert.y", true)
	v.SetDefault("calibration.coefficient.x", 0.7)
	v.SetDefault("calibration.coefficient.y", 0.5)
	v.SetDefault("calibration.tolerate.x", 1)
	v.SetDefault("calibration.tolerate.y", 1)
}

// toConfig loads the viper key-value pairs into Config.
func toConfig(v *viper.Viper) (*Config, error) {
	var cfg Config

	// unmarshal
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal into the configuration: %w", err)
	}

	return &cfg, nil
}
