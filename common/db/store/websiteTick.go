package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WebsiteStatus int

var statusName = map[WebsiteStatus]string{
	Up:      "up",
	Down:    "down",
	Unknown: "unknown",
}

const (
	Up WebsiteStatus = iota
	Down
	Unknown
)

type WebsiteTick struct {
	ID             string        `json:"id"`
	Time           time.Time     `json:"time"`
	ResponseTimeMS int64         `json:"response_time_ms"`
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

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, t := range ticks {
		queryCtx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		_, err := tx.Exec(queryCtx, query, t.Time, t.ResponseTimeMS, statusName[t.Status], t.RegionID, t.WebsiteID)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
