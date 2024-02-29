package main

import (
	"surena/node/database"
	_ "surena/node/env"
	"surena/node/scheduler"
	"surena/node/utils"
)

func main() {
	logger := utils.CreateLogger("main")

	_, err := database.Initialize()
	if err != nil {
		logger.Panic("Failed to initialize database")
	}

	sch, err := scheduler.Initialize()
	if err != nil {
		logger.Panic("Failed to initialize scheduler")
	}

	sch.Start()
	defer sch.Stop()
}
