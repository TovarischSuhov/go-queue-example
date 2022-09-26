package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/TovarischSuhov/go-queue-example/internal/api"
	"github.com/TovarischSuhov/go-queue-example/internal/kafka"
	"github.com/TovarischSuhov/go-queue-example/internal/server"
)

func main() {
	reader := kafka.NewReader("topic-1")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}
		var m api.Message
		err = json.Unmarshal(msg.Value, &m)
		if err != nil {
			log.Println(err)
			continue
		}
		go server.UpdateCallback(m)
		err = reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
