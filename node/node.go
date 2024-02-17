package main

import (
	_ "surena/node/database"
	"surena/node/scheduler"
	"surena/node/server"
	"surena/node/xray"
)

func main() {
	defer xray.Get().GetCore().Stop()
	defer scheduler.Get().Stop()
	defer server.Get().Stop()

	server.Get().Start()
}
