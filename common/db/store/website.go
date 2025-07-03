package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Website struct {
	ID        string        `json:"id"`
	Url       string        `json:"url"`
	Frequency time.Duration `json:"frequency"`
	Regions   []Region      `json:"regions"`
	CreatedAt time.Time     `json:"created_at"`
	CreatedBy string        `json:"created_by"`
}

type WebsiteStorage struct {
	db *pgxpool.Pool
}

func (s *WebsiteStorage) CreateWebsite(ctx context.Context, w Website, userId string) (*string, error) {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	websiteQuery := `
			INSERT INTO "website" (url, frequency, created_by)
			VALUES ($1, $2, $3)
			RETURNING id
	`
	queryCtx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err = tx.QueryRow(queryCtx, websiteQuery, w.Url, w.Frequency, userId).Scan(&w.ID)
	if err != nil {
		return nil, err
	}

	regionQuery := `
		INSERT INTO "website_region" (website_id, region_id)
		VALUES ($1, $2)	
	`

	for _, region := range w.Regions {
		queryCtx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		_, err = tx.Exec(queryCtx, regionQuery, w.ID, region.ID)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &w.ID, nil
}

func (s *WebsiteStorage) GetWebsiteById(ctx context.Context, id string, userId string) (*Website, error) {
	query := `
		SELECT 
            w.id,
            w.url,
            w.frequency,
            w.created_at,
            w.created_by,
            r.id,
            r.name
        FROM 
            website w
        LEFT JOIN 
            website_region wr ON w.id = wr.website_id
        LEFT JOIN 
            region r ON wr.region_id = r.id
        WHERE 
            w.id = $1 AND w.created_by = $2
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.Query(ctx, query, id, userId)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	defer rows.Close()

	website := &Website{}

	for rows.Next() {
		var region Region

		err = rows.Scan(
			&website.ID,
			&website.Url,
			&website.Frequency,
			&website.CreatedAt,
			&website.CreatedBy,
			&region.ID,
			&region.Name,
		)
		if err != nil {
			return nil, err
		}

		website.Regions = append(website.Regions, region)
	}

	if website.ID == "" {
		return nil, ErrNotFound
	}

	return website, nil
}

func (s *WebsiteStorage) GetWebsiteByFrequency(ctx context.Context, freq string) ([]Website, error) {
	query := `
		SELECT 
            w.id,
            w.url,
            w.frequency,
            w.created_at,
			w.created_by,
            r.id,
            r.name
        FROM 
            website w
        LEFT JOIN 
            website_region wr ON w.id = wr.website_id
        LEFT JOIN 
            region r ON wr.region_id = r.id
        WHERE 
            w.frequency = $1
	 `

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.Query(ctx, query, freq)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	defer rows.Close()

	var websites []Website

	for rows.Next() {
		var w Website
		var r Region

		err = rows.Scan(
			&w.ID,
			&w.Url,
			&w.Frequency,
			&w.CreatedAt,
			&w.CreatedBy,
			&r.ID,
			&r.Name,
		)
		if err != nil {
			return nil, err
		}

		w.Regions = append(w.Regions, r)

		websites = append(websites, w)
	}

	return websites, nil
}
