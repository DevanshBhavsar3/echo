package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DevanshBhavsar3/echo/api/internal/types"
	"github.com/DevanshBhavsar3/echo/api/pkg"
	"github.com/DevanshBhavsar3/echo/common/db/store"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type WebsiteHandler struct {
	websiteStorage store.WebsiteStorage
	regionStorage  store.RegionStorage
	tickStorage    store.WebsiteTickStorage
}

func NewWebsiteHandler(websiteStorage store.WebsiteStorage, regionStorage store.RegionStorage, tickStorage store.WebsiteTickStorage) *WebsiteHandler {
	return &WebsiteHandler{
		websiteStorage,
		regionStorage,
		tickStorage,
	}
}

func (h *WebsiteHandler) AddWebsite(c *fiber.Ctx) error {
	var body types.AddWebsiteBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body.",
		})
	}

	if err := pkg.Validate.Struct(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body.",
		})
	}

	userID := c.Locals("userID").(string)

	freq, err := time.ParseDuration(body.Frequency)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid freq.",
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
					"error": "Invalid region provided.",
				})
			default:
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to get regions.",
				})
			}
		}

		newWebsite.Regions = append(newWebsite.Regions, *region)
	}

	id, err := h.websiteStorage.CreateWebsite(c.Context(), newWebsite, userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating website.",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func (h *WebsiteHandler) GetAllWebsites(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	websites, err := h.websiteStorage.GetAllWebsites(c.Context(), userID)
	if err != nil && err != store.ErrNotFound {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting websites.",
		})
	}

	var response types.GetAllWebsitesResponse

	for _, w := range websites {
		ticks, err := h.tickStorage.GetLatestStatus(c.Context(), w.ID)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error getting statues for website.",
			})
		}

		website := types.WebsiteWithTicks{
			ID:        w.ID,
			Url:       w.Url,
			Frequency: pkg.ShortDuration(w.Frequency),
			CreatedAt: w.CreatedAt.Format(time.RFC3339),
			Ticks:     ticks,
		}

		for _, r := range w.Regions {
			website.Regions = append(website.Regions, r.Name)
		}

		response = append(response, website)
	}

	return c.Status(http.StatusOK).JSON(response)
}

func (h *WebsiteHandler) GetWebsiteById(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	websiteId := c.Params("id")

	err := uuid.Validate(websiteId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid website id.",
		})
	}

	website, err := h.websiteStorage.GetWebsiteById(c.Context(), websiteId, userID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Website not found.",
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error getting website.",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(website)
}

func (h *WebsiteHandler) DeleteWebsite(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	websiteId := c.Params("id")

	err := uuid.Validate(websiteId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid website id.",
		})
	}

	err = h.websiteStorage.DeleteWebsite(c.Context(), websiteId, userID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Website not found.",
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error deleting website.",
			})
		}
	}

	return c.SendStatus(http.StatusNoContent)
}

func (h *WebsiteHandler) UpdateWebsite(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	websiteId := c.Params("id")

	var body types.UpdateWebsiteBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body.",
		})
	}

	if err := pkg.Validate.Struct(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid body.",
		})
	}

	freq, err := time.ParseDuration(body.Frequency)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid frequency.",
		})
	}

	updatedWebsite := store.Website{
		ID:        websiteId,
		Url:       body.Url,
		Frequency: freq,
	}
	for _, r := range body.Regions {
		region, err := h.regionStorage.GetRegionByName(c.Context(), r)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid region provided.",
				})
			default:
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to get regions.",
				})
			}
		}

		updatedWebsite.Regions = append(updatedWebsite.Regions, *region)
	}

	err = h.websiteStorage.UpdateWebsite(c.Context(), updatedWebsite, userID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Website not found.",
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error updating website.",
			})
		}
	}

	return c.SendStatus(http.StatusNoContent)
}

func (h *WebsiteHandler) GetTicks(c *fiber.Ctx) error {
	websiteId := c.Params("id")

	err := uuid.Validate(websiteId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid website id.",
		})
	}

	start, err := time.Parse(time.RFC3339, c.Query("start"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid time range.",
		})
	}

	end, err := time.Parse(time.RFC3339, c.Query("end"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid time range.",
		})
	}

	ticks, err := h.tickStorage.GetTicks(c.Context(), websiteId, start, end)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "No ticks found for this website.",
			})
		default:
			fmt.Println("Error getting ticks:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error getting ticks.",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(ticks)
}
