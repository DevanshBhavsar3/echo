package handler

import (
	"net/http"

	"github.com/DevanshBhavsar3/common/store"
	"github.com/DevanshBhavsar3/echo-api/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WebsiteHandler struct {
	app *shared.Application
}

func NewWebsiteHandler(app *shared.Application) *WebsiteHandler {
	return &WebsiteHandler{
		app: app,
	}
}

type AddWebsitePayload struct {
	Url       string `json:"url" validate:"url"`
	Frequency string `json:"frequency" validate:"oneof=30sec 1min 3min 5min"`
}

func (h *WebsiteHandler) AddWebsite(c *fiber.Ctx) error {
	var body AddWebsitePayload

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

	newWebsite := store.Website{
		Url:       body.Url,
		Frequency: body.Frequency,
	}

	id, err := h.app.Store.Website.CreateWebsite(c.Context(), newWebsite)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "error creating website.",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func (h *WebsiteHandler) GetWebsiteById(c *fiber.Ctx) error {
	websiteId := c.Params("id")

	err := uuid.Validate(websiteId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid website id.",
		})
	}

	website, err := h.app.Store.Website.GetWebsiteById(c.Context(), websiteId)
	if err != nil {
		if err == store.ErrNotFound {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "website not found.",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "error getting website.",
		})
	}

	return c.Status(http.StatusOK).JSON(website)
}
