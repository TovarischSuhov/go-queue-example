package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/TovarischSuhov/go-queue-example/internal/client"
	"github.com/TovarischSuhov/go-queue-example/internal/kafka"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go serve()
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("Example%d", i)
		sleep := i % 3
		log.Printf("Send message '%s'\n", name)
		client.SendMessage(name, sleep, i)
	}
	wg.Wait()
}

func serve() {
	r := kafka.NewReader("topic-2")
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Finish: '%s'", msg.Value)
		err = r.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
