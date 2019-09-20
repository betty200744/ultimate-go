package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
)

type ConsumerGroupHandler struct {
}

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("topic: %s, patitions: %d, msg: %s, offset: %d \n", msg.Topic, msg.Partition, msg.Value, msg.Offset)
		session.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true

	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	consumerGroup, errC := sarama.NewConsumerGroupFromClient("group_test13", client)
	if errC != nil {
		fmt.Println("this is consumerGroup error: ", errC.Error())
	}
	defer consumerGroup.Close()
	handle := new(ConsumerGroupHandler)

	// track errors
	go func() {
		for err := range consumerGroup.Errors() {
			fmt.Println("group ERROR: ", err.Error())
		}
	}()

	for {
		err := consumerGroup.Consume(context.Background(), []string{"test13"}, handle)
		if err != nil {
			fmt.Println("consume Error: ", err.Error())
		}
	}
}
