package server

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"surena/node/server/controllers"
)

var server *Server

type Server struct {
	Host string
	Port int
	App  *fiber.App
}

func New(host string, port int) *Server {
	app := fiber.New()
	app.Mount("/", controllers.NewMainController())

	server = &Server{
		Host: host,
		Port: port,
		App:  app,
	}

	return server
}

func Get() *Server {
	return server
}

func (s *Server) GetAddress() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

func (s *Server) Start() {
	err := s.App.Listen(s.GetAddress())
	if err != nil {
		panic(err)
	}
}

func (s *Server) Stop() {
	err := s.App.Shutdown()
	if err != nil {
		return
	}
}
