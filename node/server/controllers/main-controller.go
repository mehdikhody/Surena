package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type MainController struct {
	app *fiber.App
}

func NewMainController(module *fiber.App) *MainController {
	controller := &MainController{
		app: fiber.New(),
	}

	controller.app.Get("/", controller.home)
	module.Mount("/", controller.app)
	return controller
}

func (c *MainController) home(ctx *fiber.Ctx) error {
	return ctx.SendString(fmt.Sprintf("CPU: %v", 1))
}
