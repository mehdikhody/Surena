package server

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"surena/node/server/controllers"
	"surena/node/utils"
)

var server *Server

type Server struct {
	ServerInterface
	logger         *zap.SugaredLogger
	host           string
	port           int
	app            *fiber.App
	isRunning      bool
	mainController *controllers.MainController
}

type ServerInterface interface {
	Start()
	Stop()
	GetMainController() *controllers.MainController
}

func init() {
	logger, err := utils.NewLogger("scheduler")
	if err != nil {
		fmt.Println("failed to create logger for server")
		return
	}

	app := fiber.New()
	server = &Server{
		logger:         logger,
		host:           utils.GetServerHost(),
		port:           utils.GetServerPort(),
		app:            app,
		isRunning:      false,
		mainController: controllers.NewMainController(app),
	}
}

func Get() (*Server, error) {
	if server == nil {
		return nil, errors.New("server is not initialized")
	}

	return server, nil
}

func (s *Server) Start() {
	if s.isRunning {
		s.logger.Warn("server is already running")
		return
	}

	address := fmt.Sprintf("%s:%d", s.host, s.port)
	err := s.app.Listen(address)
	if err != nil {
		s.logger.Error("failed to start server", err)
		return
	}

	s.isRunning = true
}

func (s *Server) Stop() {
	if !s.isRunning {
		s.logger.Warn("server is not running")
		return
	}

	err := s.app.Shutdown()
	if err != nil {
		s.logger.Error("failed to shutdown server", err)
		return
	}

	s.isRunning = false
}
