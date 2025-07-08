package db

import (
	"context"
	"log"

	"github.com/DevanshBhavsar3/echo/common/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context) *pgxpool.Pool {
	DATABASE_URL := config.Get("DATABASE_URL")

	dbPool, err := pgxpool.New(ctx, DATABASE_URL)
	if err != nil {
		log.Fatalf("Unable to connect to database:\n%v", err)
	}

	if err := dbPool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database:\n%v", err)
	}

	return dbPool
}
