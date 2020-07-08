package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8"
	"os"
	"time"
)

var ctx = context.Background()

func init() {
	os.Setenv("ELASTIC_APM_SERVICE_NAME", "go-redis")
	os.Setenv("ELASTIC_APM_SERVER_URL", "http://127.0.0.1:8200")
}
func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rdb.AddHook(apmgoredis.NewHook())
	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(pong, err)
	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, _ := rdb.Get(ctx, "key").Result()
	fmt.Println("test set key", val)
	time.Sleep(time.Second * 5)
}
