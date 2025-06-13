package main

import (
	"github.com/DevanshBhavsar3/echo-api/handler/v1"
	"github.com/DevanshBhavsar3/echo-api/shared"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Website interface {
		AddWebsite(c *fiber.Ctx) error
		GetWebsiteById(c *fiber.Ctx) error
	}
	Auth interface {
		Register(c *fiber.Ctx) error
		SignIn(c *fiber.Ctx) error
		Logout(c *fiber.Ctx) error
	}
}

func NewHandler(app *shared.Application) Handler {
	return Handler{
		Website: handler.NewWebsiteHandler(app),
		Auth:    handler.NewAuthHandler(app)}
}
