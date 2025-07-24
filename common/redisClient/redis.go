package redisClient

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/DevanshBhavsar3/echo/common/config"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

var WebsiteStream = "echo:websites"
var DatabaseStream = "echo:ticks"

func NewRedisClient(ctx context.Context) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:        config.Get("REDIS_URL"),
		MaxRetries:  3,
		DialTimeout: time.Second * 5,
	})

	for i := 0; i < 3; i++ {
		_, err := client.Ping(ctx).Result()
		if err == nil {
			break
		}

		if i == 2 {
			log.Fatalf("failed connecting to redis:\n%v", err)
		}

		time.Sleep(time.Second * 5)
	}

	return RedisClient{
		Client: client,
	}
}

func (r RedisClient) XAdd(ctx context.Context, stream string, data any) error {
	_, err := r.Client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]any{
			"data": data,
		},
	}).Result()

	if err != nil {
		log.Printf("failed to add data to redis stream:\n%v", err)
		return err
	}

	return nil
}

func (r RedisClient) XRead(ctx context.Context, stream string) []redis.XStream {
	res, err := r.Client.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, "$"},
		Count:   100,
	}).Result()

	if err != nil {
		if err == redis.Nil {
			return []redis.XStream{}
		}

		log.Printf("failed to read from stream:\n%v", err)
		return []redis.XStream{}
	}

	return res
}

func (r RedisClient) XGroupCreate(ctx context.Context, stream string, region string) {
	_, err := r.Client.XGroupCreateMkStream(ctx, stream, region, "$").Result()

	if err != nil && !strings.Contains(err.Error(), "exists") {
		log.Fatalf("error creating group:\n%v", err)
	}
}

func (r RedisClient) XReadGroup(ctx context.Context, stream string, group string, consumer string) []redis.XStream {
	res, err := r.Client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Streams:  []string{stream, ">"},
		Group:    group,
		Consumer: consumer,
		Count:    10,
		Block:    time.Second * 5,
	}).Result()

	if err != nil {
		if err == redis.Nil {
			fmt.Println("Redis stream is empty.")
			return []redis.XStream{}
		}

		log.Printf("failed to read from stream:\n%v", err)
		return []redis.XStream{}
	}

	return res
}

func (r RedisClient) XAck(ctx context.Context, stream string, group string, ids ...string) {
	_, err := r.Client.XAck(ctx, stream, group, ids...).Result()

	if err != nil {
		log.Printf("error acknowledging messages:\n%v", err)
	}
}
