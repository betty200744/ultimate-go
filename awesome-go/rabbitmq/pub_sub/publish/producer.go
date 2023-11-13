package main

import (
	"ultimate-go/awesome-go/rabbitmq"
	"ultimate-go/utils"
)

func main() {
	exchange := "amp.fanout"
	mq := rabbitmq.New()
	ch, err := mq.GetChannel()
	if err != nil {
		panic(err)
	}
	err = ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil)
	if err != nil {
		return
	}
	for i := 0; i < 5; i++ {
		mq.Send(ch, exchange, "", "pub_sub", utils.RandomString(8))
	}
}
