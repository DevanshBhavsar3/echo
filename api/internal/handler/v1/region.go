package handler

import (
	"errors"
	"net/http"

	"github.com/DevanshBhavsar3/echo/api/pkg"
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

	return c.JSON(fiber.Map{
		"regions": regions,
	})
}

func (h *RegionHandler) CreateRegion(c *fiber.Ctx) error {
	var region struct {
		Code string `json:"code" validate:"iso3166_1_alpha2"`
	}

	if err := c.BodyParser(&region); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body.",
		})
	}

	if err := pkg.Validate.Struct(region); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid region code.",
		})
	}

	if err := h.regionStorage.AddRegion(c.Context(), region.Code); err != nil {
		if errors.Is(err, store.ErrDuplicateRegion) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Region already exists.",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create region.",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Region created successfully.",
	})
}
