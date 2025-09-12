FROM golang:1.25-alpine AS build-stage

WORKDIR /app

COPY go.* ./
COPY ./worker/go.* ./worker/
COPY ./common ./common

RUN go work edit -dropuse="./api" -dropuse="./publisher" -dropuse="./db-worker" go.work

RUN cd ./worker && go mod download

COPY ./worker ./worker

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./worker/bin/echo-worker ./worker/cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/worker/bin/echo-worker ./echo-worker

RUN chmod +x /app/echo-worker

ENTRYPOINT ["/app/echo-worker"]
