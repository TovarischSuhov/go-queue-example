package server

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/TovarischSuhov/go-queue-example/internal/api"
	"github.com/TovarischSuhov/go-queue-example/internal/kafka"
	raw_kfk "github.com/segmentio/kafka-go"
)

func UpdateCallback(msg api.Message) {
	time.Sleep(time.Duration(msg.Sleep) * time.Second)
	resp := api.Response{Message: msg.Message, ID: msg.ID}
	buf, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	w := kafka.NewWriter("topic-2")
	err = w.WriteMessages(context.Background(), raw_kfk.Message{Value: buf})
	if err != nil {
		log.Println(err)
		return
	}
}
