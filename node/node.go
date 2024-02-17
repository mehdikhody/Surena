package main

import (
	_ "surena/node/database"
	_ "surena/node/scheduler"
	"surena/node/server"
	_ "surena/node/xray"
	"time"
)

func main() {
	// Wait for other services to initialize
	// and then start the server.
	time.Sleep(2 * time.Second)

	server.Get().Start()
}
