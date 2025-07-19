FROM golang:1.24-alpine AS build-stage

WORKDIR /app

COPY go.* ./
COPY ./db-worker/go.* ./db-worker/
COPY ./common ./common

RUN cd ./db-worker && go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./db-worker/bin/echo-db-worker ./db-worker/cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/db-worker/bin/echo-db-worker ./echo-db-worker

RUN chmod +x /app/echo-db-worker

ENTRYPOINT ["/app/echo-db-worker"]
