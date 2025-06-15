package store

import (
	"context"
	"database/sql"
	"time"
)

type Website struct {
	ID        string    `json:"id"`
	Url       string    `json:"url"`
	Frequency string    `json:"frequency"`
	CreatedAt time.Time `json:"created_at"`
}

type WebsiteStorage struct {
	db *sql.DB
}

func (s *WebsiteStorage) CreateWebsite(ctx context.Context, w Website) (*string, error) {
	query := `
			INSERT INTO "website" (url, frequency)
			VALUES ($1, $2)
			RETURNING id
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, w.Url, w.Frequency).Scan(&w.ID)
	if err != nil {
		return nil, err
	}

	return &w.ID, nil
}

func (s *WebsiteStorage) GetWebsiteById(ctx context.Context, id string) (*Website, error) {
	query := `
		SELECT id, url, frequency, created_at 
		FROM "website"
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	website := &Website{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&website.ID,
		&website.Url,
		&website.Frequency,
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

func (s *WebsiteStorage) GetWebsiteByFrequency(ctx context.Context, freq string) ([]Website, error) {
	query := `
		SELECT id, url, frequency, created_at 
		FROM "website"
		WHERE frequency = $1
	 `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, freq)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	websites := []Website{}
	for rows.Next() {
		var w Website

		err := rows.Scan(
			&w.ID,
			&w.Url,
			&w.Frequency,
			&w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		websites = append(websites, w)
	}

	return websites, nil
}
