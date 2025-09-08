FROM golang:1.24-alpine AS build-stage

WORKDIR /app

COPY go.* ./
COPY ./publisher/go.* ./publisher/
COPY ./common ./common

RUN go work edit -dropuse="./api" -dropuse="./db-worker" -dropuse="./worker" go.work

RUN cd ./publisher && go mod download

COPY ./publisher ./publisher

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./publisher/bin/echo-publisher ./publisher/cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/publisher/bin/echo-publisher ./echo-publisher

RUN chmod +x /app/echo-publisher

ENTRYPOINT ["/app/echo-publisher"]
