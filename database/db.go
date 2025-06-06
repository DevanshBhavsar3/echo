package database

import (
	"database/sql"
	"fmt"
	db "github.com/DevanshBhavsar3/echo/database/sqlc"
)
import _ "github.com/lib/pq"

func NewDatabase(connStr string) (*db.Queries, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to Postgres.")

	queries := db.New(conn)

	return queries, err
}
