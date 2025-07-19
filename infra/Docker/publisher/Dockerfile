FROM golang:1.24-alpine AS build-stage

WORKDIR /app

COPY go.* ./
COPY ./publisher/go.* ./publisher/
COPY ./common ./common

RUN cd ./publisher && go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./publisher/bin/echo-publisher ./publisher/cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/publisher/bin/echo-publisher ./echo-publisher

RUN chmod +x /app/echo-publisher

ENTRYPOINT ["/app/echo-publisher"]
