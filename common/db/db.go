package db

import (
	"context"
	"errors"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func New(ctx context.Context, addr string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(ctx, addr)
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

func Migrate(addr string) error {
	m, err := migrate.New("file://../common/db/migrations", addr)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
