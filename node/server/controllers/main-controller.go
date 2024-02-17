package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"surena/node/database"
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
	client, err := database.Get().GetClientModel().Create("test")
	if err != nil {
		return ctx.SendString(fmt.Sprintf("Error: %v", err))
	}

	return ctx.SendString(fmt.Sprintf("Client ID:", client.ID))
}
