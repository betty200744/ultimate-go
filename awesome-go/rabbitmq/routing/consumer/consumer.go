package main

import (
	"ultimate-go/awesome-go/rabbitmq"
)

func main() {
	exchange := "amp.direct"
	mq := rabbitmq.New()
	ch, err := mq.GetChannel()
	if err != nil {
		panic(err)
	}
	err = ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		return
	}
	go mq.Consumer(ch, exchange, "review", "oplog", "review")
	go mq.Consumer(ch, exchange, "activity", "oplog", "activity")
	go mq.Consumer(ch, exchange, "vip", "oplog", "vip")
	for {

	}
}
