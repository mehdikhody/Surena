package main

import (
	"surena/node/database"
	_ "surena/node/env"
	"surena/node/scheduler"
	"surena/node/server"
	"surena/node/utils"
	"surena/node/xray"
)

func main() {
	logger := utils.CreateLogger("main")

	db, err := database.Initialize()
	if err != nil {
		logger.Panic("Failed to initialize database")
	}

	defer db.Close()

	xr, err := xray.Initialize()
	if err != nil {
		logger.Panic("Failed to initialize xray")
	}

	defer xr.GetCore().Stop()
	xr.GetCore().Start()

	sch, err := scheduler.Initialize()
	if err != nil {
		logger.Panic("Failed to initialize scheduler")
	}

	defer sch.Stop()
	sch.Start()

	ser, err := server.Initialize()
	if err != nil {
		logger.Panic("Failed to initialize server")
	}

	defer ser.Stop()
	ser.Start()
}
