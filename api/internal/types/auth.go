package types

type RegisterUserBody struct {
	Name     string `json:"name" validate:"min=3,max=30"`
	Email    string `json:"email" validate:"email,max=255"`
	Avatar   string `json:"avatar" validate:"url"`
	Password string `json:"password" validate:"min=3,max=72"`
}

type SignInBody struct {
	Email    string `json:"email" validate:"email,max=255"`
	Password string `json:"password" validate:"min=3,max=72"`
}
