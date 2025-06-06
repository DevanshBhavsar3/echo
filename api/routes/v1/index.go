package v1

import (
	"github.com/DevanshBhavsar3/echo/api/handlers"
	db "github.com/DevanshBhavsar3/echo/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

func V1Router(app *fiber.App, queries *db.Queries) {
	v1Router := app.Group("/api/v1")

	websiteHandler := handlers.NewWebsiteHandler(queries)

	WebsiteRouter(v1Router, websiteHandler)
}
