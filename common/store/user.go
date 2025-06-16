package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Avatar      string    `json:"avatar"`
	Password    password  `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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

type UserStore struct {
	db *pgxpool.Pool
}

func (s *UserStore) Create(ctx context.Context, u User) (*string, error) {
	query := `
		INSERT INTO "user" (first_name, last_name, email, phone_number, avatar, password)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRow(ctx, query, u.FirstName, u.LastName, u.Email, u.PhoneNumber, u.Avatar, u.Password.hash).Scan(&u.ID)

	if err != nil {
		fmt.Println(err)
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "user_email_key" (SQLSTATE 23505)`:
			return nil, ErrDuplicateEmail
		case err.Error() == `ERROR: duplicate key value violates unique constraint "user_phone_number_key" (SQLSTATE 23505)`:
			return nil, ErrDuplicatePhoneNumber
		default:
			return nil, err
		}
	}

	return &u.ID, nil
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, first_name, last_name, email, phone_number, avatar, password, created_at
		FROM "user"
		WHERE email = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.Avatar,
		&user.Password.hash,
		&user.CreatedAt,
	)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (s *UserStore) GetById(ctx context.Context, id string) (*User, error) {
	query := `
		SELECT id, first_name, last_name, email, phone_number, avatar, password, created_at
		FROM "user"
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &User{}
	err := s.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.Avatar,
		&user.Password.hash,
		&user.CreatedAt,
	)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
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
