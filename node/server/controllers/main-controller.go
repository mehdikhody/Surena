package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type MainController struct {
	MainControllerInterface
	Controller *fiber.App
	App        *fiber.App
}

type MainControllerInterface interface {
	home(ctx *fiber.Ctx) error
}

func NewMainController(app *fiber.App) MainControllerInterface {
	c := &MainController{
		App:        app,
		Controller: fiber.New(),
	}

	c.Controller.Get("/", c.home)
	c.App.Mount("/", c.Controller)
	return c
}

func (c *MainController) home(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World!")
}
