package main

import (
	"net/http"

	"github.com/DevanshBhavsar3/echo-api/handler/v1"
	"github.com/DevanshBhavsar3/echo-api/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func mountApp(a *shared.Application) *fiber.App {
	app := fiber.New()

	app.Use(logger.New())

	// Create route handlers
	handlers := NewHandler(a)

	middlwares := handler.NewMiddlewares(a)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("I'm up.")
	})

	// v1 Router
	v1 := app.Group("/api/v1")

	// auth route
	authRouter := v1.Group("/auth")
	authRouter.Post("/register", handlers.Auth.Register)
	authRouter.Post("/signin", handlers.Auth.SignIn)
	authRouter.Post("/logout", handlers.Auth.Logout)

	// website route
	websiteRouter := v1.Group("/website", middlwares.AuthMiddleware)
	websiteRouter.Post("/", handlers.Website.AddWebsite)
	websiteRouter.Get("/:id", handlers.Website.GetWebsiteById)

	return app
}

func run(a *shared.Application, app *fiber.App) error {
	err := app.Listen(a.Port)
	return err
}
