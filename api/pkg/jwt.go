package pkg

import (
	"fmt"
	"time"

	"github.com/DevanshBhavsar3/echo/common/config"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_SECRET string
	Exp        time.Duration
	Iss        string
)

type JWTPayload struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Image   string `json:"image"`
	IsAdmin bool   `json:"is_admin"`
}

func init() {
	JWT_SECRET = config.Get("JWT_SECRET")
	Exp = time.Hour * 3
	Iss = "echo-api"
}

func GenerateJWT(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedJWT, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	return signedJWT, nil
}

func ValidateJWT(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(JWT_SECRET), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(Iss),
		jwt.WithIssuer(Iss),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
