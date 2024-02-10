package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"surena/node/scheduler"
)

var controller *MainController

type MainController struct {
	app *fiber.App
}

func NewMainController() *MainController {
	if controller != nil {
		return controller
	}

	controller := &MainController{
		app: fiber.New(),
	}

	controller.app.Get("/", controller.home)
	return controller
}

func (c *MainController) GetApp() *fiber.App {
	return c.app
}

func (c *MainController) home(ctx *fiber.Ctx) error {
	htop := scheduler.Get().HtopTask

	return ctx.SendString(fmt.Sprintf("CPU: %v", htop.CPU))
}
