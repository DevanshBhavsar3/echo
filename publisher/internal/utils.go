package internal

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/DevanshBhavsar3/echo/common/db/store"
	"github.com/DevanshBhavsar3/echo/common/redisClient"
)

func StartInterval(ctx context.Context, storage store.Storage, client redisClient.RedisClient, name string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Printf("Starting interval %s with duration %v", name, interval)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			AddWebsite(ctx, storage, client, name)
		}
	}
}

func AddWebsite(ctx context.Context, storage store.Storage, client redisClient.RedisClient, freq string) {
	payload, err := storage.Website.GetWebsiteByFrequency(ctx, freq)
	if err != nil {
		log.Printf("Failed to get websites data:\n%v", err)
		return
	}

	log.Printf("Publishing %d websites for frequency %s", len(payload), freq)

	for _, w := range payload {
		data, err := json.Marshal(w)
		if err != nil {
			log.Printf("failed to marshal website:\n%v", err)
			continue
		}

		err = client.XAdd(ctx, redisClient.WebsiteStream, data)
		if err != nil {
			log.Printf("failed to add website to stream:\n%v", err)
			continue
		}
	}
}
