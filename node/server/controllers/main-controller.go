package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"surena/node/scheduler"
)

type MainController struct {
	App *fiber.App
}

func NewMainController() *fiber.App {
	app := fiber.New()
	controller := &MainController{
		App: app,
	}

	controller.RegisterRoutes()
	return app
}

func (c *MainController) RegisterRoutes() {
	c.App.Get("/", c.Home)
}

func (c *MainController) Home(ctx *fiber.Ctx) error {
	htop := scheduler.GetScheduler().HtopTask

	return ctx.SendString(fmt.Sprintf("CPU: %v", htop.CPU))
}
