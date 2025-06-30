package main

import (
	"context"
	"log"
	"net/http"

	"github.com/DevanshBhavsar3/common/db"
	"github.com/DevanshBhavsar3/echo-api/config"
	"github.com/DevanshBhavsar3/echo-api/internal/handler/v1"
	"github.com/DevanshBhavsar3/echo-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()
	defer ctx.Done()

	PORT := config.GetEnv("PORT", ":3000")
	DATABASE_URL := config.GetEnv("DATABASE_URL", "postgres://postgres:secret@localhost:5432?sslmode=disable")

	// Connect to database
	db, err := db.New(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal("failed connecting to postgres")
	}
	defer db.Close()

	app := fiber.New()

	// Create route handlers
	handlers := handler.NewHandler(db)

	// Setup routes
	routes.SetupRoutes(app, handlers)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("I'm up.")
	})

	app.Listen(PORT)
}
