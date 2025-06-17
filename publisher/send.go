package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/DevanshBhavsar3/common"
	"github.com/DevanshBhavsar3/common/db"
	"github.com/DevanshBhavsar3/common/store"
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
		log.Fatalf("Failed to open channel.")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"websites",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange.")
	}

	forever := make(chan bool)

	go func() {
		for range time.Tick(time.Second * 3) {
			fmt.Println("Publishing 30s freq every 30 seconds.")
			websites, err := storage.Website.GetWebsiteByFrequency(ctx, "30s")
			if err != nil {
				log.Fatalf("Can't query database for website.")
			}

			// TODO: Perform batch publish
			for _, w := range websites {
				for _, r := range w.Regions {
					ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
					defer cancel()

					body, err := json.Marshal(w)
					if err != nil {
						log.Fatalf("Failed to parse website struct to json in producer.")
					}

					err = ch.PublishWithContext(
						ctx,
						"websites",
						r.Name,
						false,
						false,
						amqp.Publishing{
							DeliveryMode: amqp.Persistent,
							ContentType:  "application/json",
							Body:         []byte(body),
						})
					if err != nil {
						log.Fatal("Failed to publish message.")
					}

					fmt.Printf("Published: %v to %v queue.\n", w, r)
				}
			}

		}
	}()

	log.Printf(" [*] Publishing messages. To exit press CTRL+C")
	<-forever
}
