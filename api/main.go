package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"

	"github.com/DevanshBhavsar3/common"
	"github.com/DevanshBhavsar3/common/db"
	"github.com/DevanshBhavsar3/common/store"
	"github.com/DevanshBhavsar3/echo-api/auth"
	"github.com/DevanshBhavsar3/echo-api/handler/v1"
	"github.com/DevanshBhavsar3/echo-api/shared"
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

	ctx := context.Background()
	defer ctx.Done()

	// Connect to database
	db, err := db.New(ctx, cfg.Db.Addr)
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
	appConfig := &shared.Application{
		Config:        cfg,
		Store:         store,
		Authenticator: authenticator,
		Validator:     validator,
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
	}))

	// Create route handlers
	handlers := NewHandler(appConfig)

	middlwares := handler.NewMiddlewares(appConfig)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("I'm up.")
	})

	// v1 Router
	v1 := app.Group("/api/v1")

	// auth route
	authRouter := v1.Group("/auth")
	authRouter.Post("/register", handlers.Auth.Register)
	authRouter.Post("/signin", handlers.Auth.SignIn)
	authRouter.Post("/logout", handlers.Auth.Logout)

	// website route
	websiteRouter := v1.Group("/website", middlwares.AuthMiddleware)
	websiteRouter.Post("/", handlers.Website.AddWebsite)
	websiteRouter.Get("/:id", handlers.Website.GetWebsiteById)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server...")
		db.Close()
		os.Exit(0)
	}()

	log.Fatal(app.Listen(appConfig.Port))
}
