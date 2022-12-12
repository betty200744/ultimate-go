package main

import (
	"fmt"

	"ultimate-go/awesome-go/rabbitmq"
)

func main() {
	mq := rabbitmq.New()
	ch, err := mq.GetChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	_, err = ch.QueueDeclare("simple", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := ch.Consume("simple", "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for msg := range msgs {
			fmt.Println(fmt.Sprintf("exchange: %v, routing_key: %v, msg: %v", msg.Exchange, msg.RoutingKey, string(msg.Body[:])))
		}
	}()
	for {

	}
}
