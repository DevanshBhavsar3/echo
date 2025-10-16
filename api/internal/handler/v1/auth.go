package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/DevanshBhavsar3/echo/api/internal/types"
	"github.com/DevanshBhavsar3/echo/api/pkg"
	"github.com/DevanshBhavsar3/echo/common/config"
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

	user := store.User{
		Name:  body.Name,
		Email: body.Email,
		Image: body.Image,
	}

	// hash the user password
	if err := user.Password.Set(body.Password); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot hash the password.",
		})
	}

	newUser, err := h.userStorage.Create(c.Context(), user, "email")
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
		"sub": pkg.JWTPayload{
			ID:    newUser.ID,
			Name:  newUser.Name,
			Email: newUser.Email,
			Image: newUser.Image,
		},
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

	if string(user.Password.Hash) == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "User registered with different method.",
		})
	}

	if err := user.Password.Compare(body.Password); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password.",
		})
	}

	// Create token
	claims := jwt.MapClaims{
		"sub": pkg.JWTPayload{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Image: user.Image,
		},
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
	})
}

func (h *AuthHandler) OAuthLogin(c *fiber.Ctx) error {
	provider := c.Params("provider")

	if _, ok := pkg.OAuthConfig[provider]; !ok {
		return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/login")
	}

	state, err := pkg.GenerateRandomState()
	if err != nil {
		return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/login")
	}

	c.Cookie(
		&fiber.Cookie{
			Name:    "oauth_state",
			Value:   state,
			Expires: time.Now().Add(time.Minute * 5),
		},
	)

	url := pkg.OAuthConfig[provider].AuthCodeURL(state)

	return c.Status(http.StatusSeeOther).Redirect(url)
}

// TODO: Prettify the Error messages
func (h *AuthHandler) OAuthCallback(c *fiber.Ctx) error {
	provider := c.Params("provider")

	if _, ok := pkg.OAuthConfig[provider]; !ok {
		log.Print("Invalid provider")
		return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/login")
	}

	userState := c.Cookies("oauth_state")
	providerState := c.Query("state")

	if providerState != userState {
		log.Print("Invalid state")
		return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/login")
	}

	providerConfig := pkg.OAuthConfig[provider]

	code := c.Query("code")
	token, err := providerConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Invalid code: %v", err)
		return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/login")
	}

	oauthUser, err := providerConfig.GetOAuthUser(token)
	if err != nil {
		log.Printf("Invalid user data: %v", err)
		return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/login")
	}

	user, err := h.userStorage.GetByEmail(c.Context(), oauthUser.Email)
	if err != nil {
		switch {
		// Create a user if not found
		case errors.Is(err, store.ErrNotFound):
			user, err = h.userStorage.Create(c.Context(), *oauthUser, provider)
			if err != nil {
				return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/login")
			}
		default:
			log.Printf("Error getting user by email: %v", err)
			return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/login")
		}
	}

	// Create token
	claims := jwt.MapClaims{
		"sub": pkg.JWTPayload{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Image: user.Image,
		},
		"exp": time.Now().Add(pkg.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": pkg.Iss,
		"aud": pkg.Iss,
	}

	jwtToken, err := pkg.GenerateJWT(claims)
	if err != nil {
		return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/login")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: jwtToken,
	})

	return c.Status(http.StatusSeeOther).Redirect(config.Get("FRONTEND_URL") + "/dashboard/monitors")
}

func (h *AuthHandler) GetUser(c *fiber.Ctx) error {
	user := c.Locals("user").(pkg.JWTPayload)

	userData, err := h.userStorage.GetById(c.Context(), user.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "User not found.",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Can't get user data.",
		})
	}

	return c.Status(http.StatusOK).JSON(*userData)
}
