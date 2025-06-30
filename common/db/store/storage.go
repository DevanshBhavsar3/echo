package store

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	QueryTimeoutDuration = time.Second * 5
	ErrNotFound          = errors.New("resource not found")
)

type Storage struct {
	Website     WebsiteStorage
	Region      RegionStorage
	User        UserStorage
	WebsiteTick WebsiteTickStorage
}

func NewStorage(db *pgxpool.Pool) Storage {
	return Storage{
		Website:     WebsiteStorage{db},
		Region:      RegionStorage{db},
		User:        UserStorage{db},
		WebsiteTick: WebsiteTickStorage{db},
	}
}
