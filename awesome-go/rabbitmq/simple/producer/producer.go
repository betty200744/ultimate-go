package main

import (
	"fmt"

	"github.com/streadway/amqp"

	"gobyexample/awesome-go/rabbitmq"
	"gobyexample/utils"
)

func main() {
	mq := rabbitmq.New()
	ch, err := mq.GetChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare("simple", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		Body: []byte(utils.RandomString(8)),
	})
	if err != nil {
		fmt.Println(err)
	}
}
