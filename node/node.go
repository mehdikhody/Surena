package main

import (
	"github.com/joho/godotenv"
	"surena/node/scheduler"
	"surena/node/server"
	"surena/node/services"
)

type Node struct {
	scheduler *scheduler.Scheduler
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	logger, err := services.NewLogger("node")
	if err != nil {
		panic(err)
	}

	logger.Debug("Starting node")
	server.Initialize().Start()
}

//
//func (n *Node) StartScheduler() {
//	n.scheduler = scheduler.Get()
//	err := n.scheduler.Start()
//	if err != nil {
//		panic(err)
//	}
//}
