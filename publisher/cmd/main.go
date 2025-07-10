package main

import (
	"context"
	"time"

	"github.com/DevanshBhavsar3/echo/common/db"
	"github.com/DevanshBhavsar3/echo/common/db/store"
	"github.com/DevanshBhavsar3/echo/common/redisClient"
	"github.com/DevanshBhavsar3/echo/publisher/internal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rclient := redisClient.NewRedisClient(ctx)

	database := db.New(ctx)
	defer database.Close()

	storage := store.NewStorage(database)

	forever := make(chan bool)

	go internal.StartInterval(ctx, storage, rclient, "30s", time.Second*30)

	go internal.StartInterval(ctx, storage, rclient, "1m", time.Minute)

	go internal.StartInterval(ctx, storage, rclient, "3m", time.Minute*3)

	go internal.StartInterval(ctx, storage, rclient, "5m", time.Minute*5)

	<-forever
}
