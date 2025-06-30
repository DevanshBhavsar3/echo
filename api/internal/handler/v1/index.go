package handler

import (
	"github.com/DevanshBhavsar3/common/db/store"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Website interface {
		AddWebsite(c *fiber.Ctx) error
		GetWebsiteById(c *fiber.Ctx) error
	}
	Auth interface {
		Register(c *fiber.Ctx) error
		SignIn(c *fiber.Ctx) error
		Logout(c *fiber.Ctx) error
	}
}

func NewHandler(db *pgxpool.Pool) Handler {
	store := store.NewStorage(db)

	return Handler{
		Website: NewWebsiteHandler(store.Website, store.Region),
		Auth:    NewAuthHandler(store.User)}
}
