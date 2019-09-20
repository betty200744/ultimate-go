package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	producer, _ := sarama.NewSyncProducer([]string{"localhost:9092", "localhost:9093"}, nil)

	msg := &sarama.ProducerMessage{
		Topic: "test13",
		Key:   sarama.StringEncoder([]byte(`key`)),
		Value: sarama.StringEncoder([]byte(`bbb`)),
	}
	p, o, e := producer.SendMessage(msg)
	if e != nil {
		log.Fatal(e.Error())
	}
	fmt.Printf("partition: %d, offset: %d", p, o)
}
