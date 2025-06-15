package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/DevanshBhavsar3/common"
	"github.com/DevanshBhavsar3/common/db"
	"github.com/DevanshBhavsar3/common/store"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := db.New(common.GetEnv("DATABASE_URL", "postgres://postgres:secret@localhost:5432?sslmode=disable"))
	if err != nil {
		log.Fatal("Failed to connect to postgres.")
	}
	store := store.NewStorage(db)

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

	forever := make(chan bool)

	go func() {
		for range time.Tick(time.Second * 30) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			websites, err := store.Website.GetWebsiteByFrequency(ctx, "30sec")
			if err != nil {
				log.Fatalf("Can't query database for website.")
			}

			for _, w := range websites {
				body, err := json.Marshal(w)
				if err != nil {
					log.Fatalf("Failed to parse website struct to json in producer.")
				}

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()

				err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "application/json",
					Body:         []byte(body),
				})
				if err != nil {
					log.Fatal("Failed to publish message.")
				}
			}

		}
	}()

	log.Printf(" [*] Publishing messages. To exit press CTRL+C")
	<-forever
}
