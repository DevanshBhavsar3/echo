package shared

import (
	"time"

	"github.com/DevanshBhavsar3/common/store"
	"github.com/DevanshBhavsar3/echo-api/auth"
	"github.com/go-playground/validator/v10"
)

type Application struct {
	Config
	Store         store.Storage
	Authenticator auth.Authenticator
	Validator     *validator.Validate
}

type Config struct {
	Port string
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Addr string
}

type AuthConfig struct {
	Secret string
	Exp    time.Duration
	Iss    string
}
