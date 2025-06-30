package routes

import (
	"github.com/DevanshBhavsar3/echo-api/internal/handler/v1"
	"github.com/DevanshBhavsar3/echo-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, handlers handler.Handler) {
	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
	}))

	// v1 Routes
	v1Router := app.Group("/api/v1")

	// Auth routes
	authRouter := v1Router.Group("/auth")
	authRouter.Post("/register", handlers.Auth.Register)
	authRouter.Post("/signin", handlers.Auth.SignIn)
	authRouter.Post("/logout", handlers.Auth.Logout)

	// Website routes
	websiteRouter := v1Router.Group("/website", middleware.AuthMiddleware)
	websiteRouter.Post("/", handlers.Website.AddWebsite)
	websiteRouter.Get("/:id", handlers.Website.GetWebsiteById)
}
