package db

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DATABASE_URL string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading database .env:\n%v", err)
	}

	value, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		DATABASE_URL = "postgres://postgres:secret@localhost:5431?sslmode=disable"
	}

	DATABASE_URL = value
}

func New(ctx context.Context) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(ctx, DATABASE_URL)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := dbPool.Ping(ctx); err != nil {
		return nil, err
	}

	return dbPool, nil
}
