package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DevanshBhavsar3/common/store"
	"github.com/DevanshBhavsar3/echo-api/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	app *shared.Application
}

func NewAuthHandler(app *shared.Application) *AuthHandler {
	return &AuthHandler{
		app: app,
	}
}

type RegisterUserBody struct {
	FirstName   string `json:"first_name" validate:"min=2,max=10"`
	LastName    string `json:"last_name" validate:"min=2,max=10"`
	Email       string `json:"email" validate:"email,max=255"`
	PhoneNumber string `json:"phone_number" validate:"e164"`
	Avatar      string `json:"avatar" validate:"url"`
	Password    string `json:"password" validate:"min=3,max=72"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body RegisterUserBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse body",
		})
	}

	if err := h.app.Validator.Struct(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	user := &store.User{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
		Avatar:      body.Avatar,
	}

	// hash the user password
	if err := user.Password.Set(body.Password); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "can not hash the password",
		})
	}

	id, err := h.app.Store.User.Create(c.Context(), *user)
	if err != nil {
		if err == store.ErrDuplicateEmail {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "user already exists.",
			})
		}

		if err == store.ErrDuplicatePhoneNumber {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "phone number already used.",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "error creating user.",
		})
	}

	// Create token
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(h.app.Config.Auth.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": h.app.Config.Auth.Iss,
		"aud": h.app.Config.Auth.Iss,
	}

	token, err := h.app.Authenticator.GenerateJWT(claims)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "can not create token",
		})
	}

	cookie := fiber.Cookie{Name: "token", Value: token, Expires: time.Now().Add(h.app.Config.Auth.Exp)}
	c.Cookie(&cookie)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"msg": "registered in succcessfully.",
	})
}

type CreateTokenBody struct {
	Email    string `json:"email" validate:"email,max=255"`
	Password string `json:"password" validate:"min=3,max=72"`
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	var body CreateTokenBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	if err := h.app.Validator.Struct(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	user, err := h.app.Store.User.GetByEmail(c.Context(), body.Email)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "user not exists.",
			})
		default:
			fmt.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to sign in",
			})
		}
	}

	if err := user.Password.Compare(body.Password); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized.",
		})
	}

	// Create token
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(h.app.Config.Auth.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": h.app.Config.Auth.Iss,
		"aud": h.app.Config.Auth.Iss,
	}

	token, err := h.app.Authenticator.GenerateJWT(claims)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "can not create token",
		})
	}

	cookie := fiber.Cookie{Name: "token", Value: token, Expires: time.Now().Add(h.app.Config.Auth.Exp)}
	c.Cookie(&cookie)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"msg": "signed in succcessfully.",
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{Name: "token", Expires: time.Now()}

	c.Cookie(&cookie)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"msg": "signed out successfully.",
	})
}
