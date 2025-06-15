package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/DevanshBhavsar3/common"
	"github.com/DevanshBhavsar3/common/db"
	"github.com/DevanshBhavsar3/common/store"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx := context.Background()
	defer ctx.Done()

	db, err := db.New(ctx, common.GetEnv("DATABASE_URL", "postgres://postgres:secret@localhost:5432?sslmode=disable"))
	if err != nil {
		log.Fatal("Failed to connect to postgres.")
	}
	defer db.Close()

	storage := store.NewStorage(db)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect the queue.")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"websites-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue.")
	}

	msg, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare consumer.")
	}

	forever := make(chan bool)

	go func() {
		ctx := context.Background()
		defer ctx.Done()

		for d := range msg {
			var website store.Website

			log.Printf("Received a message: %s", d.Body)

			err := json.Unmarshal(d.Body, &website)
			if err != nil {
				log.Fatalf("Failed to parse website struct to json in consumer.")
			}

			res, err := Ping(website.Url)
			if err != nil {
				// TODO: Add to the queue back
				_ = fmt.Errorf("Failed to ping website: %v", res)
				continue
			}

			// TODO: Complete this
			ticks := []store.WebsiteTick{
				{
					Time:           time.Now(),
					ResponseTimeMS: 00,
					Status:         store.Up,
					RegionID:       "",
					WebsiteID:      website.ID,
				},
			}

			storage.WebsiteTick.BatchInsertTicks(ctx, ticks)
			fmt.Println(res)

			log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

var client = http.Client{
	Transport: &http.Transport{
		Dial: (&net.Dialer{Timeout: 2 * time.Second}).Dial,
	},
}

func Ping(url string) (*http.Response, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	return res, nil
}
