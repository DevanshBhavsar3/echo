package middleware

import (
	"net/http"
	"strings"

	"github.com/DevanshBhavsar3/echo/api/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")

	if authorizationHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header not provided.",
		})
	}

	token := strings.Split(authorizationHeader, "Bearer ")[1]

	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token not provided.",
		})
	}

	// Verify token
	jwtToken, err := pkg.ValidateJWT(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token.",
		})
	}

	claims, _ := jwtToken.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)

	// Add userID to request body
	c.Locals("userID", userID)
	return c.Next()
}
