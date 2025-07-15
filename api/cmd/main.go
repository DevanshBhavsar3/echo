package main

import (
	"context"
	"log"

	"github.com/DevanshBhavsar3/echo/api/internal/handler/v1"
	"github.com/DevanshBhavsar3/echo/api/internal/routes"
	"github.com/DevanshBhavsar3/echo/common/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()
	defer ctx.Done()

	// Connect to database
	database := db.New(ctx)
	defer database.Close()

	app := fiber.New()

	// Create route handlers
	handlers := handler.NewHandler(database)

	// Setup routes
	routes.SetupRoutes(app, handlers)

	log.Fatal(app.Listen(":3001"))
}
