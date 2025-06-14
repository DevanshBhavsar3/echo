package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthenticator struct {
	secret string
	iss    string
	aud    string
}

func NewJWTAuthenticator(secret, iss, aud string) *JWTAuthenticator {
	return &JWTAuthenticator{
		secret,
		iss,
		aud,
	}
}

func (a *JWTAuthenticator) GenerateJWT(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedJWT, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}

	return signedJWT, nil
}

func (a *JWTAuthenticator) ValidateJWT(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(a.secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(a.aud),
		jwt.WithIssuer(a.iss),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
