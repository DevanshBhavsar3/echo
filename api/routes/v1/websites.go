package v1

import (
	"github.com/DevanshBhavsar3/echo/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func WebsiteRouter(app fiber.Router, handler handlers.WebsiteHandler) {
	websitesRouter := app.Group("/websites")

	websitesRouter.Post("/", handler.AddWebsite)
	websitesRouter.Get("/:id", handler.GetWebsiteById)
}
