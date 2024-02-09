package main

import (
	"flag"
	"fmt"
	"os"
	"surena/node/database"
	"surena/node/scheduler"
	"surena/node/server"
	"surena/node/xray"
)

type Node struct {
	db        *database.Database
	xray      *xray.Xray
	scheduler *scheduler.Scheduler
	server    *server.Server
}

func main() {
	node := &Node{}
	node.StartDatabase()
	node.StartXray()
	node.StartTasks()
	node.StartServer()
}

func (n *Node) StartDatabase() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filepath := fmt.Sprintf("%s/../.bin/node-1.db", cwd)
	n.db = database.New(filepath)
}

func (n *Node) StartXray() {
	n.xray = xray.New()
}

func (n *Node) StartTasks() {
	timezone := flag.String("timezone", "UTC", "Timezone")
	flag.Parse()

	fmt.Println("Timezone:", *timezone)

	n.scheduler = scheduler.New(*timezone)
	n.scheduler.Start()
}

func (n *Node) StartServer() {
	n.server = server.New("localhost", 3000)
	n.server.Start()
}
