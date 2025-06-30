package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Region struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (r *Region) ValidateAndGetID(region string) {
	// set region id here
}

type RegionStorage struct {
	db *pgxpool.Pool
}

func (s *RegionStorage) GetAllRegions(ctx context.Context) ([]Region, error) {
	query := `
		SELECT id, name
		FROM "region"	
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var regions []Region
	for rows.Next() {
		var r Region

		err := rows.Scan(
			&r.ID,
			&r.Name,
		)
		if err != nil {
			return nil, err
		}

		regions = append(regions, r)
	}

	return regions, nil
}

func (s *RegionStorage) AddRegion(ctx context.Context, name string) error {
	query := `
		INSERT INTO "region" (name)
		VALUES ($1)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.Query(ctx, query, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *RegionStorage) GetRegionByName(ctx context.Context, name string) (*Region, error) {
	query := `
		SELECT id, name
		FROM "region"	
		WHERE name = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var region Region
	err := s.db.QueryRow(ctx, query, name).Scan(
		&region.ID,
		&region.Name,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &region, nil
}
