package types

import "github.com/DevanshBhavsar3/echo/common/db/store"

type AddWebsiteBody struct {
	Url       string   `json:"url" validate:"url"`
	Frequency string   `json:"frequency" validate:"oneof=30s 1m 3m 5m"`
	Regions   []string `json:"regions" validate:"min=1,dive,iso3166_1_alpha2"`
}

type AddWebsiteResponse struct {
	Id string `json:"id"`
}

type WebsiteWithTicks struct {
	ID        string         `json:"id"`
	Url       string         `json:"url"`
	Frequency string         `json:"frequency"`
	Regions   []string       `json:"regions"`
	CreatedAt string         `json:"created_at"`
	Ticks     []store.Status `json:"ticks"`
}

type GetAllWebsitesResponse = []WebsiteWithTicks
