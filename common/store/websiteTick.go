package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WebsiteStatus int

const (
	Up WebsiteStatus = iota
	Down
	Unknown
)

type WebsiteTick struct {
	ID             string        `json:"id"`
	Time           time.Time     `json:"time"`
	ResponseTimeMS time.Duration `json:"response_time_ms"`
	Status         WebsiteStatus `json:"status"`
	RegionID       string        `json:"region_id"`
	WebsiteID      string        `json:"website_id"`
	Website        Website       `json:"website"`
	Region         Region        `json:"region"`
}

type WebsiteTickStorage struct {
	db *pgxpool.Pool
}

// TODO: Use batch insertion
func (s *WebsiteTickStorage) BatchInsertTicks(ctx context.Context, ticks []WebsiteTick) error {
	query := `
		INSERT INTO "website_tick" (time, response_time_ms, status, region_id, website_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	batch := &pgx.Batch{}

	for _, t := range ticks {
		batch.Queue(query, t.Time, t.ResponseTimeMS, t.Status, t.RegionID, t.WebsiteID)
	}

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	br := s.db.SendBatch(ctx, batch)
	defer br.Close()

	return nil
}
