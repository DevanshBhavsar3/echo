package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/DevanshBhavsar3/echo/common/config"
	"github.com/DevanshBhavsar3/echo/common/db"
	"github.com/DevanshBhavsar3/echo/common/db/store"
	"github.com/DevanshBhavsar3/echo/worker/pkg"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var stream = "echo:websites"

var (
	REGION = config.Get("REGION")
)

func main() {
	fmt.Println(REGION, "worker started.")

	database := db.New(ctx)
	defer database.Close()

	storage := store.NewStorage(database)

	region, err := storage.Region.GetRegionByName(ctx, REGION)
	if err != nil {
		log.Fatalf("failed to determine region:\n%v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: config.Get("REDIS_URL"),
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed connecting to redis:\n%v", err)
	}

	_, err = client.XGroupCreateMkStream(ctx, stream, REGION, "$").Result()
	if err != nil && !strings.Contains(err.Error(), "exists") {
		log.Fatalf("error creating group:\n%v", err)
	}

	forever := make(chan bool)

	go func() {
		for {
			res, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
				Streams:  []string{stream, ">"},
				Group:    REGION,
				Consumer: config.Get("WORKER_ID"),
				Count:    10,
			}).Result()
			if err != nil {
				log.Fatalf("failed to read from stream:\n%v", err)
			}

			for _, i := range res {
				var processedMsg []string

				for _, j := range i.Messages {
					data := j.Values["website"].(string)

					var website store.Website

					err := json.Unmarshal([]byte(data), &website)
					if err != nil {
						log.Fatalf("error parsing redis data:\n%v", err)
					}

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

					analyst := pkg.NewAnalytics(website.Url)
					analyst.Ping()

					tick := store.WebsiteTick{
						Time:           time.Now(),
						ResponseTimeMS: analyst.ResponseTimeMS,
						Status:         analyst.Status,
						RegionID:       region.ID,
						WebsiteID:      website.ID,
					}

					encodedTick, err := json.Marshal(tick)
					if err != nil {
						log.Fatalf("error marshaling tick:\n%v", err)
					}

					// Add to db worker
					_, err = client.XAdd(ctx, &redis.XAddArgs{
						Stream: "echo:ticks",
						Values: map[string]interface{}{
							"tick": encodedTick,
						},
					}).Result()
					if err != nil {
						log.Fatalf("failed to add tick to redis:\n%v", err)
					}

					fmt.Println("Added tick to stream")

					processedMsg = append(processedMsg, j.ID)
				}

				if len(processedMsg) > 0 {
					_, err := client.XAck(ctx, stream, REGION, processedMsg...).Result()
					if err != nil {
						log.Fatalf("error acknowledging message:\n%v", err)
					}
				}
			}
		}
	}()

	<-forever
}
