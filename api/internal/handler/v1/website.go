package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DevanshBhavsar3/common/db/store"
	"github.com/DevanshBhavsar3/echo/api/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WebsiteHandler struct {
	websiteStorage store.WebsiteStorage
	regionStorage  store.RegionStorage
}

func NewWebsiteHandler(websiteStorage store.WebsiteStorage, regionStorage store.RegionStorage) *WebsiteHandler {
	return &WebsiteHandler{
		websiteStorage,
		regionStorage,
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

	if err := pkg.Validate.Struct(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	userID := c.Locals("userID").(string)

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
		region, err := h.regionStorage.GetRegionByName(c.Context(), r)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
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

	id, err := h.websiteStorage.CreateWebsite(c.Context(), newWebsite, userID)
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
	userID := c.Locals("userID").(string)
	websiteId := c.Params("id")

	err := uuid.Validate(websiteId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid website id.",
		})
	}

	website, err := h.websiteStorage.GetWebsiteById(c.Context(), websiteId, userID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "website not found.",
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "error getting website.",
			})
		}

	}

	return c.Status(http.StatusOK).JSON(website)
}
