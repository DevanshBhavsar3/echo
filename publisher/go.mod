module github.com/DevanshBhavsar3/echo/publisher

go 1.24.4

require (
	github.com/DevanshBhavsar3/echo/common/db v0.0.0-00010101000000-000000000000
	github.com/DevanshBhavsar3/echo/common/redisClient v0.0.0-00010101000000-000000000000
)

require (
	github.com/DevanshBhavsar3/echo/common/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.5 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/redis/go-redis/v9 v9.11.0 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/text v0.26.0 // indirect
)

replace (
	github.com/DevanshBhavsar3/echo/common/config => ../common/config
	github.com/DevanshBhavsar3/echo/common/db => ../common/db
	github.com/DevanshBhavsar3/echo/common/redisClient => ../common/redisClient
)
