package handlers

import (
	"fmt"
	db "github.com/DevanshBhavsar3/echo/database/sqlc"
	"github.com/gofiber/fiber/v2"
)

type WebsiteHandler struct {
	db *db.Queries
}

func NewWebsiteHandler(queries *db.Queries) WebsiteHandler {
	return WebsiteHandler{db: queries}
}

func (h WebsiteHandler) AddWebsite(c *fiber.Ctx) error {
	// TODO: Add data to database.
	return c.JSON(fiber.Map{
		"msg": "Not implemented.",
	})
}

func (h WebsiteHandler) GetWebsiteById(c *fiber.Ctx) error {
	websites, err := h.db.ListWebsites(c.Context())
	if err != nil {
		return err
	}

	fmt.Println(websites)

	websiteId := c.Params("id")

	return c.SendString(websiteId)
}
