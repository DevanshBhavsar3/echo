package store

import (
	"context"
	"errors"
	"fmt"
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

type MetricData struct {
	Current  string `json:"current"`
	Previous string `json:"previous,omitempty"`
}

type LatenciesMetrics struct {
	P99 MetricData `json:"P99"`
	P95 MetricData `json:"P95"`
	P90 MetricData `json:"P90"`
}

type Metrics struct {
	Response     LatenciesMetrics `json:"response"`
	Status       LatenciesMetrics `json:"status"`
	Availability LatenciesMetrics `json:"availability"`
}

type WebsiteTick struct {
	ID             *string   `json:"id,omitempty"`
	Time           time.Time `json:"time"`
	ResponseTimeMS *int64    `json:"responseTime,omitempty"`
	Status         string    `json:"status,omitempty"`
	RegionID       *string   `json:"region_id,omitempty"`
	WebsiteID      *string   `json:"website_id,omitempty"`
}

type Uptime struct {
	Time            string `json:"time"`
	Availability    string `json:"availability"`
	AvgResponseTime string `json:"avg_response_time"`
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

	var status []WebsiteTick = []WebsiteTick{}

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
			time_bucket('5 minutes', wt.time) as bucket,
			AVG(wt.response_time_ms)::numeric::integer as avg_respones_times
		FROM "website_tick" wt
		JOIN "region" r ON wt.region_id = r.id
		WHERE
			wt.website_id = $1
			AND wt.time BETWEEN NOW() - ($2::int * INTERVAL '1 day') AND NOW()
			AND r.name = $3
		GROUP BY bucket
		ORDER BY bucket ASC
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

	return ticks, nil
}

func (s *WebsiteTickStorage) GetMetrics(ctx context.Context, websiteID string, region string) (*Metrics, error) {
	// Response times
	response_time_query := `
		WITH filtered_ticks AS (
			SELECT
				wt.response_time_ms,
				wt.time
			FROM website_tick wt
			JOIN region r ON wt.region_id = r.id
			WHERE
				wt.website_id = $1
				AND wt.time BETWEEN NOW() - INTERVAL '2 months' AND NOW()
				AND r.name = $2
		),
		percentiles AS (
			SELECT
				percentile_cont(0.99) WITHIN GROUP (ORDER BY response_time_ms)
					FILTER (WHERE time >= NOW() - INTERVAL '1 month') AS curr_p99,
				percentile_cont(0.95) WITHIN GROUP (ORDER BY response_time_ms)
					FILTER (WHERE time >= NOW() - INTERVAL '1 month') AS curr_p95,
				percentile_cont(0.90) WITHIN GROUP (ORDER BY response_time_ms)
					FILTER (WHERE time >= NOW() - INTERVAL '1 month') AS curr_p90,

				percentile_cont(0.99) WITHIN GROUP (ORDER BY response_time_ms)
					FILTER (WHERE time < NOW() - INTERVAL '1 month') AS prev_p99,
				percentile_cont(0.95) WITHIN GROUP (ORDER BY response_time_ms)
					FILTER (WHERE time < NOW() - INTERVAL '1 month') AS prev_p95,
				percentile_cont(0.90) WITHIN GROUP (ORDER BY response_time_ms)
					FILTER (WHERE time < NOW() - INTERVAL '1 month') AS prev_p90
			FROM filtered_ticks
		)
		SELECT *
		FROM (
			SELECT
				COALESCE(curr_p99::numeric(12,2), 0) AS p99,
				COALESCE(curr_p95::numeric(12,2), 0) AS p95,
				COALESCE(curr_p90::numeric(12,2), 0) AS p90
			FROM percentiles
			UNION ALL
			SELECT
				COALESCE(prev_p99::numeric(12,2), 0) AS p99,
				COALESCE(prev_p95::numeric(12,2), 0) AS p95,
				COALESCE(prev_p90::numeric(12,2), 0) AS p90
			FROM percentiles
		)
	`

	response_time_ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.Query(response_time_ctx, response_time_query, websiteID, region)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	response_time_metric := LatenciesMetrics{}

	if rows.Next() {
		err := rows.Scan(&response_time_metric.P99.Current, &response_time_metric.P95.Current, &response_time_metric.P90.Current)
		if err != nil {
			return nil, err
		}
	}

	if rows.Next() {
		err := rows.Scan(&response_time_metric.P99.Previous, &response_time_metric.P95.Previous, &response_time_metric.P90.Previous)
		if err != nil {
			return nil, err
		}
	}

	// Status
	status_query := `
		WITH status_data AS (
			SELECT
				CASE wt.status
					WHEN 'up' THEN 1
					WHEN 'down' THEN 0
					ELSE NULL
				END as status_numeric
			FROM website_tick wt
			JOIN "region" r ON wt.region_id = r.id
			WHERE
				website_id = $1
				AND wt.time BETWEEN date_trunc('month', NOW()) AND date_trunc('month', NOW()) + INTERVAL '1 month'
				AND r.name = $2
		)
		SELECT
			CASE percentile_disc(0.99) WITHIN GROUP (ORDER BY status_numeric)
				WHEN 1 THEN 'Up'
				WHEN 0 THEN 'Down'
				ELSE 'Unknown'
			END AS p99_status,
			CASE percentile_disc(0.95) WITHIN GROUP (ORDER BY status_numeric)
				WHEN 1 THEN 'Up'
				WHEN 0 THEN 'Down'
				ELSE 'Unknown'
			END AS p95_status,
			CASE percentile_disc(0.90) WITHIN GROUP (ORDER BY status_numeric)
				WHEN 1 THEN 'Up'
				WHEN 0 THEN 'Down'
				ELSE 'Unknown'
			END AS p90_status
		FROM status_data
	`

	status_ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err = s.db.Query(status_ctx, status_query, websiteID, region)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	status_metric := LatenciesMetrics{}

	for rows.Next() {
		err := rows.Scan(&status_metric.P99.Current, &status_metric.P95.Current, &status_metric.P90.Current)
		if err != nil {
			return nil, err
		}
	}

	// Availability
	availability_query := `
		WITH buckets AS (
			SELECT
				time_bucket('5 minutes', wt.time) AS bucket,
				100.0 * AVG((wt.status = 'up')::int) AS availability_pct
			FROM "website_tick" wt
			JOIN "region" r ON wt.region_id = r.id
			WHERE
				wt.website_id = $1
				AND wt.time BETWEEN NOW() - INTERVAL '2 month' AND NOW()
				AND r.name = $2
			GROUP BY bucket
		),
		percentiles AS (
			SELECT
				percentile_cont(0.99) WITHIN GROUP (ORDER BY availability_pct)
					FILTER (WHERE bucket >= NOW() - INTERVAL '1 month') AS curr_p99,
				percentile_cont(0.95) WITHIN GROUP (ORDER BY availability_pct)
					FILTER (WHERE bucket >= NOW() - INTERVAL '1 month') AS curr_p95,
				percentile_cont(0.90) WITHIN GROUP (ORDER BY availability_pct)
					FILTER (WHERE bucket >= NOW() - INTERVAL '1 month') AS curr_p90,

				percentile_cont(0.99) WITHIN GROUP (ORDER BY availability_pct)
					FILTER (WHERE bucket < NOW() - INTERVAL '1 month') AS prev_p99,
				percentile_cont(0.95) WITHIN GROUP (ORDER BY availability_pct)
					FILTER (WHERE bucket < NOW() - INTERVAL '1 month') AS prev_p95,
				percentile_cont(0.90) WITHIN GROUP (ORDER BY availability_pct)
					FILTER (WHERE bucket < NOW() - INTERVAL '1 month') AS prev_p90
			FROM buckets
		)
		SELECT *
		FROM (
			SELECT
				COALESCE(curr_p99::numeric(5,2), 0) AS p99,
				COALESCE(curr_p95::numeric(5,2), 0) AS p95,
				COALESCE(curr_p90::numeric(5,2), 0) AS p90
			FROM percentiles
			UNION ALL
			SELECT
				COALESCE(prev_p99::numeric(5,2), 0) AS p99,
				COALESCE(prev_p95::numeric(5,2), 0) AS p95,
				COALESCE(prev_p90::numeric(5,2), 0) AS p90
			FROM percentiles
		)
	`

	availability_ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err = s.db.Query(availability_ctx, availability_query, websiteID, region)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	availability_metric := LatenciesMetrics{}

	if rows.Next() {
		err := rows.Scan(&availability_metric.P99.Current, &availability_metric.P95.Current, &availability_metric.P90.Current)
		if err != nil {
			return nil, err
		}
	}

	if rows.Next() {
		err := rows.Scan(&availability_metric.P99.Previous, &availability_metric.P95.Previous, &availability_metric.P90.Previous)
		if err != nil {
			return nil, err
		}
	}

	metrics := Metrics{
		Response:     response_time_metric,
		Status:       status_metric,
		Availability: availability_metric,
	}

	return &metrics, nil
}

type Range struct {
	From time.Time
	To   time.Time
}

func (s *WebsiteTickStorage) GetWebsiteUptime(ctx context.Context, websiteID string, uptime_range []Range) ([]Uptime, error) {
	query := `
		SELECT
			COALESCE(100.0 * (SUM(CASE WHEN wt.status = 'up' THEN 1 ELSE 0 END)::float / COUNT(*))::numeric(5,2), 0),
			COALESCE(AVG(wt.response_time_ms)::numeric(12,2), 0)
		FROM "website_tick" wt
		WHERE
			wt.website_id = $1
			AND wt.time BETWEEN $2 AND $3
	`

	var uptime []Uptime

	for _, r := range uptime_range {
		ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		var u Uptime

		u.Time = fmt.Sprintf("%v, %v", r.From.Format("2006-01-02"), r.To.Format("2006-01-02"))

		err := s.db.QueryRow(ctx, query, websiteID, r.From, r.To).Scan(&u.Availability, &u.AvgResponseTime)
		if err != nil {
			return nil, err
		}

		u.Availability = fmt.Sprintf("%v%%", u.Availability)
		u.AvgResponseTime = fmt.Sprintf("%v MS", u.AvgResponseTime)

		uptime = append(uptime, u)
	}

	return uptime, nil
}
