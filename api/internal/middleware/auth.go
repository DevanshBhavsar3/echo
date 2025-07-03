package middleware

import (
	"net/http"

	"github.com/DevanshBhavsar3/echo/api/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "token not provided.",
		})
	}

	// Verify token
	jwtToken, err := pkg.ValidateJWT(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token.",
		})
	}

	claims, _ := jwtToken.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)

	// Add userID to request body
	c.Locals("userID", userID)
	return c.Next()
}
