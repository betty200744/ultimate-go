package main

import (
	"fmt"

	"ultimate-go/awesome-go/rabbitmq"
	"ultimate-go/utils"
)

func main() {
	consumerName := fmt.Sprintf("consumer_%v", utils.RandInt(10000, 100000))
	mq := rabbitmq.New()
	ch, err := mq.GetChannel()
	if err != nil {
		panic(err)
	}
	msgs, err := ch.Consume("queue1", consumerName, true, false, false, false, nil)
	if err != nil {

	}
	go func() {
		for msg := range msgs {
			fmt.Println(fmt.Sprintf("consumer: %v, exchange: %v, routing_key: %v, msg: %v", consumerName, msg.Exchange, msg.RoutingKey, string(msg.Body[:])))
		}
	}()
	for {

	}
}
