include .env
MIGRAION_PATH = ../common/db/migrations

test:
	@go test $(shell go list -f '{{.Dir}}/...' -m | xargs)

migration:
	@echo "Enter migration name: "; \
	read MIGRAION_NAME; \
	migrate create -ext sql -dir=${MIGRAION_PATH} -seq $${MIGRAION_NAME}

migration-up:
	@migrate -path=${MIGRAION_PATH} -database ${DATABASE_URL} -verbose up

migration-down:
	@migrate -path=${MIGRAION_PATH} -database ${DATABASE_URL} -verbose down

migration-fix:
	@migrate -path=${MIGRAION_PATH} -database ${DATABASE_URL} force $(or $(VERSION), 1)