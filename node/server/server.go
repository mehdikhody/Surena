package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xtls/xray-core/common/errors"
	"surena/node/env"
	"surena/node/server/controllers"
	"surena/node/server/controllers/api"
	"surena/node/utils"
)

var server *Server
var logger = utils.CreateLogger("server")

type Server struct {
	ServerInterface
	Host           string
	Port           int
	App            *fiber.App
	Started        bool
	MainController controllers.MainControllerInterface
	APIController  api.APIControllerInterface
}

type ServerInterface interface {
	IsRunning() bool
	Start()
	Stop()
	GetMainController() controllers.MainControllerInterface
}

func init() {
	logger.Debug("initializing server")

	app := fiber.New()
	server = &Server{
		Host:           env.GetServerHost(),
		Port:           env.GetServerPort(),
		App:            app,
		Started:        false,
		MainController: controllers.NewMainController(app),
		APIController:  api.NewAPIController(app),
	}
}

func Initialize() (ServerInterface, error) {
	if server == nil {
		return nil, errors.New("server is not initialized")
	}

	return server, nil
}

func Get() ServerInterface {
	if server == nil {
		panic("server is not initialized")
	}

	return server
}

func (s *Server) Start() {
	if s.Started {
		logger.Warn("server is already running")
		return
	}

	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	err := s.App.Listen(address)
	if err != nil {
		logger.Error("failed to start server", err)
		return
	}

	logger.Infof("server is running on %s", address)
	s.Started = true
}

func (s *Server) Stop() {
	if !s.Started {
		logger.Warn("server is not running")
		return
	}

	err := s.App.Shutdown()
	if err != nil {
		logger.Error("failed to shutdown server", err)
		return
	}

	logger.Info("server is stopped")
	s.Started = false
}

func (s *Server) GetMainController() controllers.MainControllerInterface {
	return s.MainController
}
