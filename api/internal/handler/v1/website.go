package handler

import (
	"errors"
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
			Regions:   w.Regions,
		}

		response = append(response, website)
	}

	return c.Status(http.StatusOK).JSON(response)
}

var DefaultUptimeRanges = []store.Range{
	{
		From: time.Now().Truncate(24 * time.Hour),
		To:   time.Now(),
	},
	{
		From: time.Now().AddDate(0, 0, -7),
		To:   time.Now(),
	},
	{
		From: time.Now().AddDate(0, -1, 0),
		To:   time.Now(),
	},
	{
		From: time.Now().AddDate(-1, 0, 0),
		To:   time.Now(),
	},
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

	uptime, err := h.tickStorage.GetWebsiteUptime(c.Context(), websiteId, DefaultUptimeRanges)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting website uptime.",
		})
	}

	response := types.GetWebsiteByIdResponse{
		ID:        website.ID,
		Url:       website.Url,
		Frequency: pkg.ShortDuration(website.Frequency),
		Regions:   website.Regions,
		CreatedAt: website.CreatedAt.Format(time.RFC3339),
		Uptime:    uptime,
	}

	return c.Status(http.StatusOK).JSON(response)
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

	days := c.Query("days")
	if days == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Please provide tick range.",
		})
	}

	region := c.Query("region")
	if region == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Please provide region.",
		})
	}

	ticks, err := h.tickStorage.GetTicks(c.Context(), websiteId, days, region)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "No ticks found for this website.",
			})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error getting ticks.",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(ticks)
}

func (h *WebsiteHandler) GetMetrics(c *fiber.Ctx) error {
	websiteId := c.Params("id")

	region := c.Query("region")

	if region == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Please provide region.",
		})
	}

	err := uuid.Validate(websiteId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid website id.",
		})
	}

	metrics, err := h.tickStorage.GetMetrics(c.Context(), websiteId, region)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting metrics.",
		})
	}

	return c.Status(http.StatusOK).JSON(metrics)
}

func (h *WebsiteHandler) GetUptime(c *fiber.Ctx) error {
	websiteId := c.Params("id")

	from, err := time.Parse(time.DateOnly, c.Query("from"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format.",
		})
	}

	to, err := time.Parse(time.DateOnly, c.Query("to"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format.",
		})
	}

	uptime_range := store.Range{
		From: from.Truncate(24 * time.Hour),
		To:   to.Truncate(24 * time.Hour).Add(23 * time.Hour),
	}

	uptime, err := h.tickStorage.GetWebsiteUptime(c.Context(), websiteId, []store.Range{
		uptime_range,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error getting website uptime.",
		})
	}

	return c.Status(http.StatusOK).JSON(uptime[0])
}
