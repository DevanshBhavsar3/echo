package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	QueryTimeoutDuration = time.Second * 5
	ErrNotFound          = errors.New("resource not found")
)

type Storage struct {
	Website interface {
		CreateWebsite(ctx context.Context, w Website, userId string) (*string, error)
		GetWebsiteById(ctx context.Context, id string, userId string) (*Website, error)
		GetWebsiteByFrequency(ctx context.Context, freq string) ([]Website, error)
	}

	Region interface {
		AddRegion(ctx context.Context, name string) error
		GetAllRegions(ctx context.Context) ([]Region, error)
		GetRegionByName(ctx context.Context, name string) (*Region, error)
	}

	User interface {
		Create(ctx context.Context, u User) (*string, error)
		GetByEmail(ctx context.Context, email string) (*User, error)
		GetById(ctx context.Context, id string) (*User, error)
	}

	WebsiteTick interface {
		BatchInsertTicks(ctx context.Context, t []WebsiteTick) error
	}
}

func NewStorage(db *pgxpool.Pool) Storage {
	return Storage{
		Website:     &WebsiteStorage{db},
		Region:      &RegionStorage{db},
		User:        &UserStore{db},
		WebsiteTick: &WebsiteTickStorage{db},
	}
}
