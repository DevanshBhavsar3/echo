package middleware

import (
	"encoding/json"
	"fmt"
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
	payload := claims["sub"].(map[string]interface{})

	body, err := json.Marshal(payload)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token.",
		})
	}

	user := pkg.JWTPayload{}
	if err := json.Unmarshal(body, &user); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token.",
		})
	}

	fmt.Println(user)

	c.Locals("user", user)
	return c.Next()
}
