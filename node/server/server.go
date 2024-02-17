package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"surena/node/env"
	"surena/node/server/controllers"
	"surena/node/utils"
)

var server *Server

type Server struct {
	ServerInterface
	Logger         *logrus.Entry
	Host           string
	Port           int
	App            *fiber.App
	Started        bool
	MainController controllers.MainControllerInterface
}

type ServerInterface interface {
	IsRunning() bool
	Start()
	Stop()
	GetMainController() controllers.MainControllerInterface
}

func init() {
	logger := utils.CreateLogger("server")
	logger.Debug("initializing server")

	app := fiber.New()
	server = &Server{
		Logger:         logger,
		Host:           env.GetServerHost(),
		Port:           env.GetServerPort(),
		App:            app,
		Started:        false,
		MainController: controllers.NewMainController(app),
	}
}

func Get() ServerInterface {
	if server == nil {
		panic("server is not initialized")
	}

	return server
}

func (s *Server) Start() {
	if s.Started {
		s.Logger.Warn("server is already running")
		return
	}

	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	err := s.App.Listen(address)
	if err != nil {
		s.Logger.Error("failed to start server", err)
		return
	}

	s.Logger.Infof("server is running on %s", address)
	s.Started = true
}

func (s *Server) Stop() {
	if !s.Started {
		s.Logger.Warn("server is not running")
		return
	}

	err := s.App.Shutdown()
	if err != nil {
		s.Logger.Error("failed to shutdown server", err)
		return
	}

	s.Logger.Info("server is stopped")
	s.Started = false
}

func (s *Server) GetMainController() controllers.MainControllerInterface {
	return s.MainController
}
