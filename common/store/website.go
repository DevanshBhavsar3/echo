package store

import (
	"context"
	"database/sql"
	"time"
)

type Website struct {
	ID               string    `json:"id"`
	Url              string    `json:"url"`
	HealthCheckRoute string    `json:"health_check_route"`
	CreatedAt        time.Time `json:"created_at"`
}

type WebsiteStorage struct {
	db *sql.DB
}

func (m *WebsiteStorage) CreateWebsite(ctx context.Context, w Website) (*string, error) {
	query := `
			INSERT INTO "website" (url, health_check_route)
			VALUES ($1, $2)
			RETURNING id
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := m.db.QueryRowContext(ctx, query, w.Url, w.HealthCheckRoute).Scan(&w.ID)
	if err != nil {
		return nil, err
	}

	return &w.ID, nil
}

func (m *WebsiteStorage) GetWebsiteById(ctx context.Context, id string) (*Website, error) {
	query := `
		SELECT id, url, health_check_route, created_at 
		FROM "website"
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	website := &Website{}
	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&website.ID,
		&website.Url,
		&website.HealthCheckRoute,
		&website.CreatedAt,
	)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return website, nil
}
