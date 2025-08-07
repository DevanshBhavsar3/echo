package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/DevanshBhavsar3/echo/api/internal/types"
	"github.com/DevanshBhavsar3/echo/api/pkg"
	"github.com/DevanshBhavsar3/echo/common/db/store"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	userStorage store.UserStorage
}

func NewAuthHandler(userStorage store.UserStorage) *AuthHandler {
	return &AuthHandler{
		userStorage,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body types.RegisterUserBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body.",
		})
	}

	if err := pkg.Validate.Struct(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data.",
		})
	}

	user := &store.User{
		Name:   body.Name,
		Email:  body.Email,
		Avatar: body.Avatar,
	}

	// hash the user password
	if err := user.Password.Set(body.Password); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot hash the password.",
		})
	}

	newUser, err := h.userStorage.Create(c.Context(), *user)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrDuplicateEmail):
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Email already used.",
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error creating user.",
			})
		}
	}

	// Create token
	claims := jwt.MapClaims{
		"sub": newUser.ID,
		"exp": time.Now().Add(pkg.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": pkg.Iss,
		"aud": pkg.Iss,
	}

	token, err := pkg.GenerateJWT(claims)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create token.",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"token": token,
		"user":  newUser,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body types.LoginBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body.",
		})
	}

	if err := pkg.Validate.Struct(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data.",
		})
	}

	user, err := h.userStorage.GetByEmail(c.Context(), body.Email)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "User does not exists.",
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to login.",
			})
		}
	}

	if err := user.Password.Compare(body.Password); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password.",
		})
	}

	// Create token
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(pkg.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": pkg.Iss,
		"aud": pkg.Iss,
	}

	token, err := pkg.GenerateJWT(claims)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create token.",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{Name: "token", Expires: time.Now()}

	c.Cookie(&cookie)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"msg": "Logged out successfully.",
	})
}

func (h *AuthHandler) GetUser(c *fiber.Ctx) error {
	userId := c.Locals("userID").(string)

	user, err := h.userStorage.GetById(c.Context(), userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Can't get user data.",
		})
	}

	return c.Status(http.StatusOK).JSON(user)
}
