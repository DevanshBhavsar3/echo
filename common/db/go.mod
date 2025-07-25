module github.com/DevanshBhavsar3/echo/common/db

go 1.24.4

require (
	github.com/DevanshBhavsar3/echo/common/config v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.7.5
	golang.org/x/crypto v0.39.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/text v0.26.0 // indirect
)

replace github.com/DevanshBhavsar3/echo/common/config => ../config
