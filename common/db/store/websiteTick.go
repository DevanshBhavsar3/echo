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

var websiteStatusMap = map[string]WebsiteStatus{
	"up":      Up,
	"down":    Down,
	"unknown": Unknown,
}

func ParseWebsiteStatus(status string) (WebsiteStatus, error) {
	if s, ok := websiteStatusMap[status]; ok {
		return s, nil
	}

	return Unknown, errors.New("invalid website status")
}

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

type Tick struct {
	WebsiteTick
	Region
}

type WebsiteTick struct {
	ID             *string   `json:"id,omitempty"`
	Time           time.Time `json:"time"`
	ResponseTimeMS *int64    `json:"responseTime,omitempty"`
	Status         string    `json:"status,omitempty"`
	RegionID       *string   `json:"region_id,omitempty"`
	WebsiteID      *string   `json:"website_id,omitempty"`
}

type WebsiteTickStorage struct {
	db *pgxpool.Pool
}

func (s *WebsiteTickStorage) GetLatestStatus(ctx context.Context, websiteID string) ([]WebsiteTick, error) {
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
			return []WebsiteTick{}, nil
		default:
			return nil, err
		}
	}
	defer rows.Close()

	var status []WebsiteTick

	for rows.Next() {
		var tickTime pgtype.Timestamptz
		var tick WebsiteTick

		err := rows.Scan(&tickTime, &tick.Status)
		if err != nil {
			return nil, err
		}

		tick.Time = tickTime.Time

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

		_, err := tx.Exec(queryCtx, query, t.Time, t.ResponseTimeMS, t.Status, t.RegionID, t.WebsiteID)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (s *WebsiteTickStorage) GetTicks(ctx context.Context, websiteID string, days string, region string) ([]Tick, error) {
	query := `
		SELECT 
			time_bucket('15 minutes', wt.time) as bucket, 
			avg(wt.response_time_ms)::numeric::integer as avg_respones_times
		FROM "website_tick" wt 
		JOIN "region" r ON wt.region_id = r.id 
		WHERE 
			wt.website_id = $1
			AND wt.time >= NOW() - ($2::int * INTERVAL '1 day')
			AND wt.time <= NOW()
			AND r.name = $3
		GROUP BY bucket 
		ORDER BY bucket DESC
		LIMIT 500
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.Query(ctx, query, websiteID, days, region)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return []Tick{}, nil
		default:
			return nil, err
		}
	}
	defer rows.Close()

	var ticks []Tick

	for rows.Next() {
		var tickTime pgtype.Timestamptz
		var tick Tick

		err := rows.Scan(&tickTime, &tick.ResponseTimeMS)
		if err != nil {
			return nil, err
		}

		tick.WebsiteTick.Time = tickTime.Time

		ticks = append(ticks, tick)
	}

	slices.Reverse(ticks)

	return ticks, nil
}
