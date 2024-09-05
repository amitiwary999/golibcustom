package producer

import (
	"log"

	"github.com/amitiwary999/go-kyeue/producer"
	"github.com/amitiwary999/go-kyeue/storage"
)

func ProduceMessage() {
	queue, err := storage.NewPostgresClient("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable", 10, 10)
	if err != nil {
		log.Fatal("failed to create storage %v ", err)
	} else {
		producer := producer.NewQueueProducer(queue)
		err = producer.CreateChannel("prac-queue")
		if err != nil {
			log.Fatal("failed to created channel %v ", err)
		}

		payloadStr := "We are sending the test message"
		err = producer.Send(payloadStr, "prac-queue")
		if err != nil {
			log.Fatal("failed to send message in queue %v ", err)
		}
	}
}
