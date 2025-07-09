package types

type AddWebsiteBody struct {
	Url       string   `json:"url" validate:"url"`
	Frequency string   `json:"frequency" validate:"oneof=30s 1m 3m 5m"`
	Regions   []string `json:"regions" validate:"min=1,dive,iso3166_1_alpha3"`
}

type AddWebsiteResponse struct {
	Id string `json:"id"`
}
