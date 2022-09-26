package kafka

import "github.com/segmentio/kafka-go"

var (
	readers map[string]*kafka.Reader
	writers map[string]*kafka.Writer
)

func NewReader(topic string) *kafka.Reader {
	var r *kafka.Reader
	var ok bool
	if r, ok = readers[topic]; !ok {
		r = kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   topic,
			GroupID: "myAppID",
		})
		readers[topic] = r
	}
	return r
}

func NewWriter(topic string) *kafka.Writer {
	var w *kafka.Writer
	var ok bool
	if w, ok = writers[topic]; !ok {
		w = &kafka.Writer{
			Addr:  kafka.TCP("localhost:9092"),
			Topic: topic,
		}
		writers[topic] = w
	}
	return w
}
