package types

import "github.com/DevanshBhavsar3/echo/common/db/store"

type RegisterUserBody struct {
	Name     string `json:"name" validate:"min=3,max=30"`
	Email    string `json:"email" validate:"email,max=255"`
	Image    string `json:"image" validate:"url"`
	Password string `json:"password" validate:"min=8,max=72"`
}

type LoginBody struct {
	Email    string `json:"email" validate:"email,max=255"`
	Password string `json:"password" validate:"min=3,max=72"`
}

type AuthResponse struct {
	Token string     `json:"token"`
	User  store.User `json:"user"`
}
