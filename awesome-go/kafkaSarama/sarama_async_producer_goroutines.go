package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	wg                          sync.WaitGroup
	enqueued, successes, errors int
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ProducerLoop:
	for {
		message := &sarama.ProducerMessage{Topic: "test13", Value: sarama.StringEncoder("testing 123")}
		select {
		case producer.Input() <- message:
			time.Sleep(time.Second)
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			errors++
		}
	}()

	wg.Wait()
	log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
}
