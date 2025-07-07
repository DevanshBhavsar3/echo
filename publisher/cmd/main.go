package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/DevanshBhavsar3/echo/common/config"
	"github.com/DevanshBhavsar3/echo/common/db"
	"github.com/DevanshBhavsar3/echo/common/db/store"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var stream = "echo:websites"

func main() {
	REDIS_URL := config.Get("REDIS_URL")

	database := db.New(ctx)
	defer database.Close()

	storage := store.NewStorage(database)

	client := redis.NewClient(&redis.Options{
		Addr: REDIS_URL,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed connecting to redis:\n%v", err)
	}

	forever := make(chan bool)

	go func() {
		for range time.Tick(time.Second * 30) {
			AddWebsite(storage, client, "30s")
		}
	}()

	go func() {
		for range time.Tick(time.Minute) {
			AddWebsite(storage, client, "1m")
		}
	}()

	go func() {
		for range time.Tick(time.Minute * 3) {
			AddWebsite(storage, client, "3m")
		}
	}()

	go func() {
		for range time.Tick(time.Minute * 5) {
			AddWebsite(storage, client, "5m")
		}
	}()

	<-forever
}

func AddWebsite(store store.Storage, client *redis.Client, freq string) {
	fmt.Println("added")

	websites, err := store.Website.GetWebsiteByFrequency(ctx, freq)
	if err != nil {
		log.Fatalf("failed to get websites data:\n%v", err)
	}

	for _, w := range websites {
		data, err := json.Marshal(w)
		if err != nil {
			log.Fatalf("failed to marshal website:\n%v", err)
		}

		_, err = client.XAdd(ctx, &redis.XAddArgs{
			Stream: stream,
			Values: map[string]interface{}{
				"website": data,
			},
		}).Result()
		if err != nil {
			log.Fatalf("failed to add data to redis stream:\n%v", err)
		}
	}
}
