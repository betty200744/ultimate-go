package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	message := &sarama.ProducerMessage{Topic: "test13", Value: sarama.StringEncoder("testing 123")}

ProducerLoop:
	for {
		time.Sleep(time.Second)
		select {
		case producer.Input() <- message:
		case s := <-producer.Successes():
			fmt.Println(s.Value)
			successes++
		case err := <-producer.Errors():
			fmt.Println(err.Error())
			errors++
		case <-signals:
			break ProducerLoop
		}
	}
}
