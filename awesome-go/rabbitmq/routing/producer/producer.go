package main

import (
	"ultimate-go/awesome-go/rabbitmq"
	"ultimate-go/utils"
)

func main() {
	exchange := "amq.direct"
	mq := rabbitmq.New()
	ch, err := mq.GetChannel()
	if err != nil {
		panic(err)
	}
	err = ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		return
	}
	for i := 0; i < 5; i++ {
		mq.Send(ch, exchange, "", "review", utils.RandomString(8))
		mq.Send(ch, exchange, "", "activity", utils.RandomString(8))
		mq.Send(ch, exchange, "", "vip", utils.RandomString(8))
	}
}
