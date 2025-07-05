package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DevanshBhavsar3/echo/db"
	"github.com/DevanshBhavsar3/echo/db/store"
	"github.com/DevanshBhavsar3/echo/publisher/config"

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

	client := redis.NewClient(&redis.Options{
		Addr: cfg.REDIS_URL,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed connecting to redis:\n%v", err)
	}

	AddWebsite(storage, client, "30s")
}

func AddWebsite(store store.Storage, client *redis.Client, freq string) error {
	websites, err := store.Website.GetWebsiteByFrequency(ctx, freq)
	if err != nil {
		_ = fmt.Errorf("%v", err)
		return err
	}

	fmt.Println(websites)

	//	client.XAdd(ctx, &redis.XAddArgs{
	//		Stream: stream,
	//		Values: websites,
	//	})

	return nil
}
