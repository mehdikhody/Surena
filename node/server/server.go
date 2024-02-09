package server

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"surena/node/server/controllers"
)

type Server struct {
	Host string
	Port int
	App  *fiber.App
}

func New(host string, port int) *Server {
	app := fiber.New()
	app.Mount("/", controllers.NewMainController())

	return &Server{
		Host: host,
		Port: port,
		App:  app,
	}
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
