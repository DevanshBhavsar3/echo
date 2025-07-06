package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/DevanshBhavsar3/echo/db"
	"github.com/DevanshBhavsar3/echo/db/store"
	"github.com/DevanshBhavsar3/echo/worker/config"
	"github.com/DevanshBhavsar3/echo/worker/pkg"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var stream = "echo:websites"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading publisher .env:\n%v", err)
	}

	cfg := config.LoadEnv()

	database := db.New(ctx)
	defer database.Close()

	storage := store.NewStorage(database)

	region, err := storage.Region.GetRegionByName(ctx, cfg.REGION)
	if err != nil {
		log.Fatalf("failed to determine region:\n%v", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: cfg.REDIS_URL,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed connecting to redis:\n%v", err)
	}

	_, err = client.XGroupCreateMkStream(ctx, stream, cfg.REGION, "$").Result()
	if err != nil && !strings.Contains(err.Error(), "exists") {
		log.Fatalf("error creating group:\n%v", err)
	}

	forever := make(chan bool)

	go func() {
		for {
			res, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
				Streams:  []string{stream, ">"},
				Group:    cfg.REGION,
				Consumer: cfg.WORKER_ID,
				Count:    10,
			}).Result()
			if err != nil {
				log.Fatalf("failed to read from stream:\n%v", err)
			}

			for _, i := range res {
				var processedMsg []string

				// TODO: Only process the current region's websites
				for _, j := range i.Messages {
					data := j.Values["website"].(string)

					var website store.Website

					err := json.Unmarshal([]byte(data), &website)
					if err != nil {
						log.Fatalf("error parsing redis data:\n%v", err)
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

				_, err := client.XAck(ctx, stream, cfg.REGION, processedMsg...).Result()
				if err != nil {
					log.Fatalf("error acknowledging message:\n%v", err)
				}
			}
		}
	}()

	<-forever
}
