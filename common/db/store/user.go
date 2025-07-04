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
	Avatar    string    `json:"avatar"`
	Password  password  `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

func (p *password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(text))
}

type UserStorage struct {
	db *pgxpool.Pool
}

func (s *UserStorage) Create(ctx context.Context, u User) (*User, error) {
	query := `
		INSERT INTO "user" (name, email, avatar, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, avatar, created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var user User

	err := s.db.QueryRow(ctx, query, u.Name, u.Email, u.Avatar, u.Password.hash).Scan(&user.ID, &user.Name, &user.Email, &user.Avatar, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "user_email_key" (SQLSTATE 23505)`:
			return nil, ErrDuplicateEmail
		case err.Error() == `ERROR: duplicate key value violates unique constraint "user_phone_number_key" (SQLSTATE 23505)`:
			return nil, ErrDuplicatePhoneNumber
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserStorage) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, name, email, avatar, password, created_at, updated_at
		FROM "user"
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Avatar,
		&user.Password.hash,
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
		SELECT id, name, email, avatar, password, created_at, updated_at
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
		&user.Avatar,
		&user.Password.hash,
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
	ErrDuplicateEmail       = errors.New("a user with that email already exists")
	ErrDuplicatePhoneNumber = errors.New("a user with that phone number already exists")
)
