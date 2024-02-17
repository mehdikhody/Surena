package main

import (
	_ "surena/node/database"
	_ "surena/node/scheduler"
	"surena/node/server"
)

func main() {
	server.Get().Start()
}
