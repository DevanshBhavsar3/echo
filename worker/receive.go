package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/DevanshBhavsar3/common"
	"github.com/DevanshBhavsar3/common/db"
	"github.com/DevanshBhavsar3/common/store"
)

// NOTE: This will run on different regions
func main() {
	region, ok := os.LookupEnv("REGION")
	if region == "" || !ok {
		log.Fatal("Failed to determie the region.")
		return
	}

	ctx := context.Background()
	defer ctx.Done()

	// NOTE: This will connect to some central db
	db, err := db.New(ctx, common.GetEnv("DATABASE_URL", "postgres://postgres:secret@localhost:5432?sslmode=disable"))
	if err != nil {
		log.Fatal("Failed to connect to postgres.")
	}
	defer db.Close()

	// storage := store.NewStorage(db)

	// NOTE: This will connect to remote rebbitmq client
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

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue.")
	}

	err = ch.QueueBind(
		q.Name,
		region,
		"websites",
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to bind queue to exchange.")
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
		fmt.Printf("Listening for messages in %v queue.\n", region)
		for d := range msg {
			var website store.Website

			log.Printf("Received a message: %s", d.Body)

			err := json.Unmarshal(d.Body, &website)
			if err != nil {
				log.Fatalf("Failed to parse website struct to json in consumer.")
			}

			// TODO: Perform analytics on the request url
			// ----------------

			GetAnalytics(website.Url)

			// TODO: Complete this
			// ticks := []store.WebsiteTick{
			// 	{
			// 		Time:           time.Now(),
			// 		ResponseTimeMS: 00,       // fix
			// 		Status:         store.Up, // fix
			// 		RegionID:       "",       // fix
			// 		WebsiteID:      website.ID,
			// 	},
			// }

			// storage.WebsiteTick.BatchInsertTicks(ctx, ticks)

			log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
