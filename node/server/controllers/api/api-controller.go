package api

import "github.com/gofiber/fiber/v2"

type APIController struct {
	APIControllerInterface
	Controller *fiber.App
	App        *fiber.App
}

type APIControllerInterface interface {
}

func NewAPIController(app *fiber.App) APIControllerInterface {
	c := &APIController{
		App:        app,
		Controller: fiber.New(),
	}

	c.App.Mount("/api", c.Controller)
	return c
}
