FROM golang:1.25-alpine AS build-stage

WORKDIR /app

COPY go.* ./
COPY ./db-worker/go.* ./db-worker/
COPY ./common ./common

RUN go work edit -dropuse="./api" -dropuse="./publisher" -dropuse="./worker" go.work

RUN cd ./db-worker && go mod download

COPY ./db-worker ./db-worker

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./db-worker/bin/echo-db-worker ./db-worker/cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/db-worker/bin/echo-db-worker ./echo-db-worker

RUN chmod +x /app/echo-db-worker

ENTRYPOINT ["/app/echo-db-worker"]
