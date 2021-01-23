package main

import (
	"gobyexample/awesome-go/rabbitmq"
	"gobyexample/utils"
)

func main() {
	mq := rabbitmq.New()
	ch, err := mq.GetChannel()
	if err != nil {
		panic(err)
	}
	ch.Qos(1, 0, false)
	for i := 0; i < 10; i++ {
		mq.Send(ch, "", "queue1", "queue1", utils.RandomString(8))
	}
}
