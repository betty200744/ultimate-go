package main

import (
	"fmt"
	"github.com/coocood/freecache"
	"runtime/debug"
)

func main() {
	size := 100 * 1024 * 1024
	cache := freecache.NewCache(size)
	debug.SetGCPercent(20)
	key := []byte(`k1`)
	val := []byte(`k1 value`)
	expire := 60
	cache.Set(key, val, expire)
	k1, errGet := cache.Get(key)
	if errGet != nil {
		fmt.Println(errGet)
	}
	fmt.Println(string(k1))
	fmt.Println("entry count: ", cache.EntryCount())
	cache.Del(key)
}
