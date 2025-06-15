package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	QueryTimeoutDuration = time.Second * 5
	ErrNotFound          = errors.New("resource not found")
)

type Storage struct {
	Website interface {
		CreateWebsite(ctx context.Context, w Website) (*string, error)
		GetWebsiteById(ctx context.Context, id string) (*Website, error)
		GetWebsiteByFrequency(ctx context.Context, freq string) ([]Website, error)
	}

	User interface {
		Create(ctx context.Context, u User) (*string, error)
		GetByEmail(ctx context.Context, email string) (*User, error)
		GetById(ctx context.Context, id string) (*User, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Website: &WebsiteStorage{db},
		User:    &UserStore{db},
	}
}
