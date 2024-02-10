package server

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"surena/node/server/controllers"
)

var server *Server
var serverInitialized = false
var serverStarted = false

type Server struct {
	host string
	port int
	app  *fiber.App
}

func Initialize() *Server {
	if serverInitialized {
		panic("Server is already initialized")
	}

	app := fiber.New()
	app.Mount("/", controllers.NewMainController())

	server = &Server{
		host: GetHost(),
		port: GetPort(),
		app:  app,
	}

	serverInitialized = true
	return server
}

func Get() *Server {
	if !serverInitialized {
		panic("Server is not initialized")
	}

	return server
}

func GetHost() string {
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	return host
}

func GetPort() int {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	return portInt
}

func (s *Server) GetAddress() string {
	return s.host + ":" + strconv.Itoa(s.port)
}

func (s *Server) Start() {
	if serverStarted {
		panic("Server is already started")
	}

	err := s.app.Listen(s.GetAddress())
	if err != nil {
		panic(err)
	}

	serverStarted = true
}

func (s *Server) Stop() {
	if !serverStarted {
		return
	}

	err := s.app.Shutdown()
	if err != nil {
		return
	}

	serverStarted = false
}
