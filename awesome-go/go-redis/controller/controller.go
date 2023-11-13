package controller

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"ultimate-go/awesome-go/go-redis/redisClient"
)

func ExampleClient(ctx context.Context) {
	rdb := redisClient.NewClient()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("test set key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
