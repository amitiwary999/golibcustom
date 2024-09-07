package gokyeuerun

import (
	"fmt"
	"log"
	"time"

	"github.com/amitiwary999/go-kyeue/producer"
	"github.com/amitiwary999/go-kyeue/storage"
)

func ProduceMessage() {
	queue, err := storage.NewPostgresClient("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable", 10, 10)
	if err != nil {
		log.Fatal("failed to create storage %v ", err)
	} else {
		producer := producer.NewQueueProducer(queue)
		err = producer.CreateChannel("prac_queue")
		if err != nil {
			log.Fatal("failed to created channel %v ", err)
		}
		time.Sleep(time.Duration(10 * time.Second))
		for i := 0; i < 10; i++ {
			payloadStr := fmt.Sprintf("We are sending the test message %d", i)
			fmt.Printf("payload %v \n", payloadStr)
			err = producer.Send(payloadStr, "prac_queue")
			if err != nil {
				log.Fatal("failed to send message in queue %v ", err)
			}
		}
	}
}
