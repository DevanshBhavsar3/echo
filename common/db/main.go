package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context) *pgxpool.Pool {
	var DATABASE_URL string

	value, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		DATABASE_URL = "postgres://postgres:secret@localhost:5432?sslmode=disable"
	} else {
		DATABASE_URL = value
	}

	dbPool, err := pgxpool.New(ctx, DATABASE_URL)
	if err != nil {
		log.Fatalf("Unable to connect to database:\n%v", err)
	}

	if err := dbPool.Ping(ctx); err != nil {
		forever := make(chan bool)

		<-forever
		log.Fatalf("Unable to ping database:\n%v", err)
	}

	return dbPool
}
