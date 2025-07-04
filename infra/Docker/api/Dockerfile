FROM golang:1.21.0-alpine AS builder

WORKDIR /app

COPY ./go.work ./go.work.sum ./
COPY ./api/go.mod ./api/go.sum ./api/
COPY ./common ./common

RUN cd ./api && go mod tidy

COPY ./api ./api

RUN go build -o ./api/bin/echo-api ./api/cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache bash curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz | \
    tar -xz -C /usr/local/bin && \
    chmod +x /usr/local/bin/migrate

COPY --from=builder /app/api/bin/echo-api ./bin/echo-api
COPY ./common/db/migrations ./migrations

EXPOSE 3000

ENTRYPOINT ["sh", "-c", "migrate -path=./migrations -database ${DATABASE_URL} up && ./bin/echo-api"]
