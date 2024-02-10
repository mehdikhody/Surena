package main

import (
	"github.com/joho/godotenv"
	"surena/node/database"
	"surena/node/scheduler"
	"surena/node/server"
	"surena/node/xray"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.Initialize()
	xray.New()
	scheduler.Initialize()
	server.Initialize()
}
