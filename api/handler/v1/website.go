package handler

import (
	"fmt"
	"net/http"
	"time"

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
	Url       string   `json:"url" validate:"url"`
	Frequency string   `json:"frequency" validate:"oneof=30s 1m 3m 5m"`
	Regions   []string `json:"regions" validate:"min=1,dive,iso3166_1_alpha3"`
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

	user := c.Locals("user").(*store.User)

	freq, err := time.ParseDuration(body.Frequency)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid freq",
		})
	}

	newWebsite := store.Website{
		Url:       body.Url,
		Frequency: freq,
	}
	for _, r := range body.Regions {
		region, err := h.app.Store.Region.GetRegionByName(c.Context(), r)
		if err != nil {
			switch {
			case err == store.ErrNotFound:
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"error": "invalid region provided",
				})
			default:
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error": "failed to get regions",
				})
			}
		}

		newWebsite.Regions = append(newWebsite.Regions, *region)
	}

	id, err := h.app.Store.Website.CreateWebsite(c.Context(), newWebsite, user.ID)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "error creating website.",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func (h *WebsiteHandler) GetWebsiteById(c *fiber.Ctx) error {
	user := c.Locals("user").(*store.User)
	websiteId := c.Params("id")

	fmt.Println(user)

	err := uuid.Validate(websiteId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid website id.",
		})
	}

	website, err := h.app.Store.Website.GetWebsiteById(c.Context(), websiteId, user.ID)
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
