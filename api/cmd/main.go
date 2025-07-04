package main

import (
	"context"
	"log"

	"github.com/DevanshBhavsar3/echo/api/internal/handler/v1"
	"github.com/DevanshBhavsar3/echo/api/internal/routes"
	"github.com/DevanshBhavsar3/echo/db"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	defer ctx.Done()

	// Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading api .env:\n%v", err)
	}

	// Connect to database
	database, err := db.New(ctx)
	if err != nil {
		log.Fatalf("failed connecting to postgres:\n%v", err)
	}
	defer database.Close()

	app := fiber.New()

	// Create route handlers
	handlers := handler.NewHandler(database)

	// Setup routes
	routes.SetupRoutes(app, handlers)

	log.Fatal(app.Listen(":3000"))
}
