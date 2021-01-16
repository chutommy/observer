package main

import (
	"errors"
	"log"

	"observer/config"
	"observer/controller"
)

const robotName = "Observer"

func main() {
	// load configuration
	cfg, err := config.GetConfig(".")
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
