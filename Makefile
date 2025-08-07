-include .env
MIGRAION_PATH = ./common/db/migrations

install:
	@echo "Installing dependencies..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@export PATH="$HOME/go/bin:$PATH"

test:
	@go test $(shell go list -f '{{.Dir}}/...' -m | xargs)

lint:
	@golangci-lint run ./api/...
	@golangci-lint run ./common/config/...
	@golangci-lint run ./common/db/...
	@golangci-lint run ./common/redisClient/...
	@golangci-lint run ./publisher/...
	@golangci-lint run ./worker/...
	@golangci-lint run ./db-worker/...

format:
	@gofmt -s -w .

format-check:
	@if [ -n "$$(gofmt -l .)" ]; then \
		exit 1; \
	fi

migration:
	@echo "Enter migration name: "; \
	read MIGRAION_NAME; \
	@migrate create -ext sql -dir=${MIGRAION_PATH} -seq $${MIGRAION_NAME}

migration-up:
	@migrate -path=${MIGRAION_PATH} -database ${DATABASE_URL} -verbose up

migration-down:
	@migrate -path=${MIGRAION_PATH} -database ${DATABASE_URL} -verbose down

migration-fix:
	@migrate -path=${MIGRAION_PATH} -database ${DATABASE_URL} force $(or $(VERSION), 1)
