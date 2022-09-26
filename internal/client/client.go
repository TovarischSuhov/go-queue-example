package client

import (
	"context"
	"encoding/json"
	"log"

	"github.com/TovarischSuhov/go-queue-example/internal/api"
	"github.com/TovarischSuhov/go-queue-example/internal/kafka"
	raw_kfk "github.com/segmentio/kafka-go"
)

func SendMessage(messageName string, sleepTime int, id int) {
	msg := api.Message{Message: messageName, Sleep: sleepTime, ID: id}
	buf, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}
	w := kafka.NewWriter("topic-1")
	err = w.WriteMessages(context.Background(), raw_kfk.Message{Value: buf})
	if err != nil {
		log.Println(err)
		return
	}
}
