package main

import (
	"errors"
	"os"

	"github.com/chutommy/observer/config"
	"github.com/chutommy/observer/controller"

	"github.com/sirupsen/logrus"
)

const (
	robotName = "Observer"

	// config file
	fileName  = "settings"
	extension = "toml"
	path      = "."
)

func main() {
	log := logger()

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

func logger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return log
}
