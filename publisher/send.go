package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pong, err := client.Ping(ctx).Result()

	fmt.Println(pong, err)
}
