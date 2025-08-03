package handler

import (
	"github.com/DevanshBhavsar3/echo/common/db/store"
	"github.com/gofiber/fiber/v2"
)

type RegionHandler struct {
	regionStorage store.RegionStorage
}

func NewRegionHandler(regionStorage store.RegionStorage) *RegionHandler {
	return &RegionHandler{
		regionStorage: regionStorage,
	}
}

func (h *RegionHandler) GetRegions(c *fiber.Ctx) error {
	regions, err := h.regionStorage.GetAllRegions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch regions.",
		})
	}

	if len(regions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No regions found.",
		})
	}

	return c.JSON(fiber.Map{
		"regions": regions,
	})
}
