package main

import (
	"fmt"

	"ultimate-go/awesome-go/rabbitmq"
	"ultimate-go/utils"
)

func main() {
	mq := rabbitmq.New()
	ch, err := mq.GetChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	err = mq.Send(ch, "", "simple", "simple", utils.RandomString(8))
	if err != nil {
		fmt.Println(err)
	}
}
