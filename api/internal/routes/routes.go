package routes

import (
	"net/http"

	"github.com/DevanshBhavsar3/echo/api/internal/handler/v1"
	"github.com/DevanshBhavsar3/echo/api/internal/middleware"
	"github.com/DevanshBhavsar3/echo/common/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, handlers handler.Handler) {
	ENV := config.Get("ENV")
	FRONTEND_URL := config.Get("FRONTEND_URL")

	corsConfig := cors.Config{
		AllowCredentials: true,
	}

	if ENV == "LOCAL" {
		corsConfig.AllowOrigins = "http://localhost:5173"
	} else {
		corsConfig.AllowOrigins = FRONTEND_URL
	}

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New(corsConfig))

	// Health Route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("I'm up.")
	})

	// v1 Routes
	v1Router := app.Group("/api/v1")

	// Auth routes
	authRouter := v1Router.Group("/auth")
	authRouter.Post("/register", handlers.Auth.Register)
	authRouter.Post("/signin", handlers.Auth.SignIn)
	authRouter.Get("/user", middleware.AuthMiddleware, handlers.Auth.GetUser)
	authRouter.Post("/logout", handlers.Auth.Logout)

	// Website routes
	websiteRouter := v1Router.Group("/website", middleware.AuthMiddleware)
	websiteRouter.Post("/", handlers.Website.AddWebsite)
	websiteRouter.Get("/:id", handlers.Website.GetWebsiteById)
}
