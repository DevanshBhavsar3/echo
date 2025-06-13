package handler

import (
	"fmt"
	"net/http"

	"github.com/DevanshBhavsar3/echo-api/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Middlwares struct {
	app *shared.Application
}

func NewMiddlewares(app *shared.Application) *Middlwares {
	return &Middlwares{
		app: app,
	}
}

func (m *Middlwares) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")

	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "token not provided.",
		})
	}

	// Verify token
	jwtToken, err := m.app.Authenticator.ValidateJWT(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token.",
		})
	}

	claims, _ := jwtToken.Claims.(jwt.MapClaims)

	user, err := m.app.Store.User.GetById(c.Context(), claims["sub"].(string))
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token.",
		})
	}

	// Add userId to body
	c.Locals("user", user)
	return c.Next()
}
