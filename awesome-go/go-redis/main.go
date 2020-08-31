package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8"
	"os"
)

func init() {
	os.Setenv("ELASTIC_APM_SERVICE_NAME", "go-redis")
	os.Setenv("ELASTIC_APM_SERVER_URL", "http://localhost:8200")
}
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	for i := 0; i < 30; i++ {
		rdb.AddHook(apmgoredis.NewHook())
		rdb.Ping(context.Background())
		err := rdb.Set(context.Background(), "key", "value", 0).Err()
		fmt.Println("test set key", err)
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < 130; i++ {
		val, _ := rdb.Get(context.Background(), "key").Result()
		fmt.Println("test get key", val)
	}
	// watch 1
	ctx := context.Background()
	_ = rdb.Watch(ctx, func(tx *redis.Tx) error {
		errSet := tx.Set(ctx, "kt1", "value", 0).Err()
		if errSet != nil {
			return errSet
		}
		res, errGet := tx.Get(ctx, "kt1").Result()
		if errGet != nil {
			return errGet
		}
		fmt.Println(res)
		return nil
	}, "kt1", "kt2")

	// watch 2
	key := "mykey"
	_ = rdb.Watch(ctx, func(tx *redis.Tx) error {
		v, errGet := tx.Get(ctx, key).Int()
		if errGet != nil {
			return errGet
		}
		v = v + 1
		errSet := tx.Set(ctx, key, v, 0).Err()
		if errSet != nil {
			return errSet
		}
		return nil
	}, key)
	v, _ := rdb.Get(ctx, key).Int()
	fmt.Println("after watch", v)
}
