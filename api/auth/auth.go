package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Authenticator interface {
	GenerateJWT(claims jwt.Claims) (string, error)
	ValidateJWT(token string) (*jwt.Token, error)
}
