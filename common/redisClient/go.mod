module github.com/DevanshBhavsar3/echo/common/redisClient

go 1.24.4

require (
	github.com/DevanshBhavsar3/echo/common/config v0.0.0-00010101000000-000000000000
	github.com/redis/go-redis/v9 v9.11.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)

replace (
	github.com/DevanshBhavsar3/echo/common/config => ../config
	github.com/DevanshBhavsar3/echo/common/db => ../db
)
