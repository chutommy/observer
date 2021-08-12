package main

import (
	"errors"
	"os"

	"observer/config"
	"observer/controller"

	"github.com/sirupsen/logrus"
)

const (
	robotName = "Observer"
	fileName  = "settings"
	extension = "toml"
	path      = "."
)

func main() {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	entryConfig := log.WithFields(logrus.Fields{
		"location":       "configuration",
		"file name":      fileName,
		"file extension": extension,
		"path":           path,
	})

	entryController := log.WithFields(logrus.Fields{
		"location":   "controller",
		"robot name": robotName,
	})

	cfg, err := config.GetConfig(entryConfig, path, fileName, extension)
	if err != nil {
		if errors.Is(err, config.ErrSettingsNotFound) {
			log.Println("settings file not found, a default settings is generated and being used...")
		} else {
			log.Fatal(err)
		}
	}

	o := controller.NewObserver(robotName, entryController, cfg)
	o.LoadWork()
	o.LoadRobot()

	if err = o.Start(); err != nil {
		log.Fatal("failed to start the Observer: %w", err)
	}
}
