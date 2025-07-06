package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/DevanshBhavsar3/echo/db"
	"github.com/DevanshBhavsar3/echo/db-worker/config"
	"github.com/DevanshBhavsar3/echo/db/store"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var stream = "echo:ticks"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading publisher .env:\n%v", err)
	}

	cfg := config.LoadEnv()

	database := db.New(ctx)
	defer database.Close()

	storage := store.NewStorage(database)

	client := redis.NewClient(&redis.Options{
		Addr: cfg.REDIS_URL,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed connecting to redis:\n%v", err)
	}

	forever := make(chan bool)

	go func() {
		var ticks []store.WebsiteTick

		for {
			res, err := client.XRead(ctx, &redis.XReadArgs{
				Streams: []string{stream, "$"},
				Count:   100,
			}).Result()
			if err != nil && err != redis.Nil {
				log.Fatalf("failed to read from stream:\n%v", err)
			}

			for _, i := range res {
				for _, j := range i.Messages {
					data := j.Values["tick"].(string)

					var tick store.WebsiteTick

					err := json.Unmarshal([]byte(data), &tick)
					if err != nil {
						log.Fatalf("error parsing redis data:\n%v", err)
					}

					ticks = append(ticks, tick)
				}
			}

			if len(ticks) > 3 {
				err := storage.WebsiteTick.BatchInsertTicks(ctx, ticks)
				if err != nil {
					log.Fatalf("error inserting ticks to db:\n%v", err)
				}

				fmt.Printf("INSERTED %v TICKS", len(ticks))
				ticks = nil
			}
		}
	}()

	<-forever
}
