FROM golang:1.24 AS build-stage

WORKDIR /app

COPY . .

RUN cd ./api && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./api/bin/echo-api ./api/cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/api/bin/echo-api ./echo-api

RUN chmod +x /app/echo-api

EXPOSE 3000

ENTRYPOINT ["/app/echo-api"]
