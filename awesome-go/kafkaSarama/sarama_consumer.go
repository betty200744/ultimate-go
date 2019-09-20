package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	partitionConsume, errP := consumer.ConsumePartition("test13", 0, sarama.OffsetNewest)
	if errP != nil {
		panic(errP)
	}
	defer partitionConsume.Close()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsume.Messages():
			fmt.Printf("partition: %d, msg : %s, offset: %d \n", msg.Partition, msg.Value, msg.Offset)
		case <-signals:
			break ConsumerLoop
		}
	}
}
