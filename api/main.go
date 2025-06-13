package main

import (
	"log"
	"time"

	"github.com/DevanshBhavsar3/common"
	"github.com/DevanshBhavsar3/common/db"
	"github.com/DevanshBhavsar3/common/store"
	"github.com/DevanshBhavsar3/echo-api/auth"
	"github.com/DevanshBhavsar3/echo-api/shared"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Load the configs
	cfg := shared.Config{
		Port: common.GetEnv("PORT", ":3000"),
		Db: shared.DbConfig{
			Addr: common.GetEnv("DATABASE_URL", "postgres://postgres:secret@localhost:5432?sslmode=disable"),
		},
		Auth: shared.AuthConfig{
			Secret: common.GetEnv("JWT_SECRET", "jwt_secret"),
			Exp:    time.Minute * 5,
			Iss:    "echo-api",
		},
	}

	// Connect to database
	db, err := db.New(cfg.Db.Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create store service
	store := store.NewStorage(db)

	// Create auth service
	authenticator := auth.NewJWTAuthenticator(cfg.Auth.Secret, cfg.Auth.Iss, cfg.Auth.Iss)

	validator := validator.New()

	// Start api server
	app := &shared.Application{
		Config:        cfg,
		Store:         store,
		Authenticator: authenticator,
		Validator:     validator,
	}

	fiber := mountApp(app)

	log.Fatal(run(app, fiber))
}
