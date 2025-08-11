package store

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WebsiteStatus int

func (s WebsiteStatus) String() string {
	switch s {
	case Up:
		return "up"
	case Down:
		return "down"
	case Unknown:
		return "unknown"
	}

	return "unknown"
}

const (
	Up WebsiteStatus = iota
	Down
	Unknown
)

type Status struct {
	Time   time.Time `json:"time"`
	Status string    `json:"status"`
}

type WebsiteTick struct {
	ID             string        `json:"id"`
	Time           time.Time     `json:"time"`
	ResponseTimeMS int64         `json:"response_time_ms"`
	Status         WebsiteStatus `json:"status"`
	RegionID       string        `json:"region_id"`
	WebsiteID      string        `json:"website_id"`
}

type WebsiteTickStorage struct {
	db *pgxpool.Pool
}

func (s *WebsiteTickStorage) GetLatestTicks(ctx context.Context, websiteID string) ([]WebsiteTick, error) {
	query := `
		SELECT id, 
					 time,
					 response_time_ms,
					 status, 
					 region_id, 
					 website_id
		FROM "website_tick"
		WHERE website_id = $1
		LIMIT 5
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.Query(ctx, query, websiteID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return []WebsiteTick{}, nil
		default:
			return nil, err
		}
	}
	defer rows.Close()

	var ticks []WebsiteTick

	for rows.Next() {
		var t WebsiteTick

		err := rows.Scan(
			&t.ID,
			&t.Time,
			&t.ResponseTimeMS,
			&t.Status,
			&t.RegionID,
			&t.WebsiteID,
		)
		if err != nil {
			return nil, err
		}

		ticks = append(ticks, t)
	}

	return ticks, nil
}

func (s *WebsiteTickStorage) GetLatestStatus(ctx context.Context, websiteID string) ([]Status, error) {
	query := `
		SELECT time, status
		FROM "website_tick"
		WHERE website_id = $1
		ORDER BY time DESC
		LIMIT 5
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.Query(ctx, query, websiteID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return []Status{}, nil
		default:
			return nil, err
		}
	}
	defer rows.Close()

	var status []Status

	for rows.Next() {
		var tickTime pgtype.Timestamptz
		var tickStatus string

		err := rows.Scan(&tickTime, &tickStatus)
		if err != nil {
			return nil, err
		}

		tick := Status{
			Time:   tickTime.Time,
			Status: tickStatus,
		}

		status = append(status, tick)
	}

	slices.Reverse(status)

	return status, nil
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
	//nolint:errcheck
	defer tx.Rollback(ctx)

	for _, t := range ticks {
		queryCtx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		_, err := tx.Exec(queryCtx, query, t.Time, t.ResponseTimeMS, t.Status.String(), t.RegionID, t.WebsiteID)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
