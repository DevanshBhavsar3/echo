package middleware

import (
	"net/http"

	"github.com/DevanshBhavsar3/echo/api/pkg"
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(pkg.JWTPayload)

	if user.IsAdmin {
		return c.Next()
	}

	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"error": "You are not an Admin.",
	})
}
