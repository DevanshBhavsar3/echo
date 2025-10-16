package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Image     string    `json:"image"`
	Password  Password  `json:"password,omitzero"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Password struct {
	Text *string
	Hash []byte
}

func (p *Password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.Text = &text
	p.Hash = hash

	return nil
}

func (p *Password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.Hash, []byte(text))
}

type UserStorage struct {
	db *pgxpool.Pool
}

func (s *UserStorage) Create(ctx context.Context, u User, provider string) (*User, error) {
	user_query := `
		INSERT INTO "user" (name, email, image)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, image, created_at, updated_at;
	`

	var account_query string

	if provider == "email" {
		account_query = `
			INSERT INTO "account" (user_id, provider, password)
			VALUES ($1, $2, $3)
		`
	} else {
		account_query = `
			INSERT INTO "account" (user_id, provider)
			VALUES ($1, $2)
		`
	}

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	//nolint:errcheck
	defer tx.Rollback(ctx)

	var user User

	user_ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err = tx.QueryRow(user_ctx, user_query, u.Name, u.Email, u.Image).Scan(&user.ID, &user.Name, &user.Email, &user.Image, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "user_email_key" (SQLSTATE 23505)`:
			return nil, ErrDuplicateEmail
		default:
			return nil, err
		}
	}

	account_ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	if provider == "email" {
		_, err = tx.Exec(account_ctx, account_query, user.ID, provider, u.Password.Hash)
	} else {
		_, err = tx.Exec(account_ctx, account_query, user.ID, provider)
	}

	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStorage) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT u.id, u.name, u.email, u.image, a.password, u.created_at, u.updated_at
		FROM "user" u
		JOIN "account" a ON u.id = a.user_id
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Image,
		&user.Password.Hash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (s *UserStorage) GetById(ctx context.Context, id string) (*User, error) {
	query := `
		SELECT id, name, email, image, created_at, updated_at
		FROM "user"
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Image,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

var (
	ErrDuplicateEmail = errors.New("a user with that email already exists")
)
