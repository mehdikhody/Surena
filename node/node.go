package main

import (
	"github.com/joho/godotenv"
	"surena/node/scheduler"
	"surena/node/server"
	"surena/node/utils"
)

type Node struct {
	scheduler *scheduler.Scheduler
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	logger, err := utils.NewLogger("node")
	if err != nil {
		panic(err)
	}

	node := &Node{}
	go node.StartScheduler()

	logger.Debug("Starting node")
	server.Initialize().Start()
}

func (n *Node) StartScheduler() {
	var err error
	n.scheduler, err = scheduler.Get()
	if err != nil {
		panic(err)
	}

	err = n.scheduler.Start()
	if err != nil {
		panic(err)
	}
}
