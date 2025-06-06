package main

import (
	"log"
	"os"

	v1 "github.com/DevanshBhavsar3/echo/api/routes/v1"
	"github.com/DevanshBhavsar3/echo/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	queries, err := database.NewDatabase("postgres://postgres:secret@localhost:5432/echo?sslmode=disable")

	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
		return
	}

	app := fiber.New()
	port := os.Getenv("PORT")

	if port == "" {
		port = ":3000"
	}

	v1.V1Router(app, queries)

	log.Fatal(app.Listen(port))
}
