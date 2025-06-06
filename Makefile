include .env

migration_up:
	$(shell migrate -path database/migration/ -database ${DATABASE_URL} -verbose up)

migration_down:
	$(shell migrate -path database/migration/ -database ${DATABASE_URL} -verbose down)

migration_fix:
	$(shell migrate -path database/migration/ -database ${DATABASE_URL} force $(or $(VERSION), 1))

generate:
	cd ./database && sqlc generate