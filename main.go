package main

import (
	"errors"
	"os"

	"observer/config"
	"observer/controller"

	"github.com/sirupsen/logrus"
)

const robotName = "Observer"

// Define settings file information.
const (
	fileName  = "settings"
	extension = "toml"
	path      = "."
)

func main() {
	// create a new logger instance
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// create config entry
	entryConfig := log.WithFields(logrus.Fields{
		"location":       "configuration",
		"file name":      fileName,
		"file extension": extension,
		"path":           path,
	})

	// load configuration
	cfg, err := config.GetConfig(entryConfig, path, fileName, extension)
	if err != nil {
		if errors.Is(err, config.ErrSettingsNotFound) {
			log.Println("settings file not found, a default settings is generated and being used...")
		} else {
			log.Fatal(err)
		}
	}

	// constructs and config the new Observer
	o := controller.NewObserver(robotName, cfg)
	o.LoadWork()
	o.LoadRobot()

	// starts the observer
	err = o.Start()
	if err != nil {
		log.Fatal("failed to start the Observer: %w", err)
	}
}

// TODO
// - add logger
// - testing
// - refactor readme file
// - code cleanup (dependencies, haar cascades etc.)
// - Docker implementation
// - install sh
