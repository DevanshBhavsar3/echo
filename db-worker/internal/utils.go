package internal

import (
	"context"
	"encoding/json"
	"log"

	"github.com/DevanshBhavsar3/echo/common/db/store"
	"github.com/redis/go-redis/v9"
)

func AddToBatch(res []redis.XStream, ticks *[]store.WebsiteTick) {
	for _, i := range res {
		for _, j := range i.Messages {
			data := j.Values["data"].(string)

			var tick store.WebsiteTick

			err := json.Unmarshal([]byte(data), &tick)
			if err != nil {
				log.Printf("error parsing redis data:\n%v", err)
				continue
			}

			*ticks = append(*ticks, tick)
		}
	}
}

func ProcessBatch(ctx context.Context, storage store.Storage, ticks *[]store.WebsiteTick) {
	err := storage.WebsiteTick.BatchInsertTicks(ctx, *ticks)
	if err != nil {
		log.Printf("error inserting ticks to db:\n%v", err)
	}

	log.Printf("Inserted %d messages to database.", len(*ticks))

	*ticks = nil
}
