package gokyeuconsumer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/amitiwary999/go-kyeue/consumer"
	"github.com/amitiwary999/go-kyeue/model"
	"github.com/amitiwary999/go-kyeue/storage"
)

type MsgHandler struct{}

func (msh *MsgHandler) MessageHandler(msg model.Message) error {
	fmt.Printf("message received %v \n", msg.Id)
	return nil
}

func StartConsumer() {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)
	dbClient, err := storage.NewPostgresClient("", 10, 10)
	if err != nil {
		log.Fatal(err)
	}
	consmr := consumer.NewQueueConsumer(dbClient, "prac-queue", 1, &MsgHandler{})
	ctx := context.Background()
	consmr.Consume(ctx)
	<-gracefulShutdown
	ctx.Done()
}