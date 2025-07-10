package main

import (
	"context"
	"time"

	"github.com/DevanshBhavsar3/echo/common/db"
	"github.com/DevanshBhavsar3/echo/common/db/store"
	"github.com/DevanshBhavsar3/echo/common/redisClient"
	"github.com/DevanshBhavsar3/echo/db-worker/internal"
)

var (
	BATCH_SIZE    = 100
	BATCH_TIMEOUT = time.Second * 5
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database := db.New(ctx)
	defer database.Close()

	storage := store.NewStorage(database)

	rclient := redisClient.NewRedisClient(ctx)

	var ticks []store.WebsiteTick

	ticker := time.NewTicker(BATCH_TIMEOUT)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if len(ticks) > 0 {
				internal.ProcessBatch(ctx, storage, &ticks)
			}
		default:
			res := rclient.XRead(ctx, redisClient.DatabaseStream)
			internal.AddToBatch(res, &ticks)

			if len(ticks) > BATCH_SIZE {
				internal.ProcessBatch(ctx, storage, &ticks)
			}
		}
	}
}
