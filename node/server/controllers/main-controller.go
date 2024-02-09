package controllers

import "github.com/gofiber/fiber/v2"

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
	return ctx.SendString("Hello, World!")
}
