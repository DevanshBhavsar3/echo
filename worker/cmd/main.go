package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/DevanshBhavsar3/echo/common/config"
	"github.com/DevanshBhavsar3/echo/common/db"
	"github.com/DevanshBhavsar3/echo/common/db/store"
	"github.com/DevanshBhavsar3/echo/common/redisClient"
	"github.com/DevanshBhavsar3/echo/worker/internal"
)

var (
	REGION    = config.Get("REGION")
	WORKER_ID = config.Get("WORKER_ID")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rclient := redisClient.NewRedisClient(ctx)

	database := db.New(ctx)
	defer database.Close()

	storage := store.NewStorage(database)

	region, err := storage.Region.GetRegionByName(ctx, REGION)
	if err != nil {
		log.Fatalf("failed to determine region:\n%v", err)
	}

	// Create consumer group
	rclient.XGroupCreate(ctx, redisClient.WebsiteStream, REGION)

	for {
		// Get messages from streams
		res := rclient.XReadGroup(ctx, redisClient.WebsiteStream, REGION, WORKER_ID)

		for _, i := range res {
			var processedMsg []string

			for _, j := range i.Messages {
				data := j.Values["data"].(string)

				var website store.Website
				err := json.Unmarshal([]byte(data), &website)
				if err != nil {
					log.Printf("error parsing redis message:\n%v", err)
					continue
				}

				// Check if the website is of this worker's region
				var supportedRegion bool
				for _, r := range website.Regions {
					if r.Name == REGION {
						supportedRegion = true
						break
					}
				}

				if !supportedRegion {
					continue
				}

				// Ping the website
				status, responseTime := internal.Ping(website.Url)

				tick := store.WebsiteTick{
					Time:           time.Now(),
					ResponseTimeMS: &responseTime,
					Status:         status.String(),
					RegionID:       region.ID,
					WebsiteID:      &website.ID,
				}

				encodedTick, err := json.Marshal(tick)
				if err != nil {
					log.Printf("error marshaling tick:\n%v", err)
					continue
				}

				// Add the tick to database stream
				err = rclient.XAdd(ctx, redisClient.DatabaseStream, encodedTick)

				if err == nil {
					log.Printf("Processed message for region %s", REGION)
					processedMsg = append(processedMsg, j.ID)
				}
			}

			// Acknowlege processed messages
			if len(processedMsg) > 0 {
				rclient.XAck(ctx, redisClient.WebsiteStream, REGION, processedMsg...)
			}
		}
	}
}
