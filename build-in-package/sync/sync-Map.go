package main

import (
	"fmt"
	"sync"
)

var smap sync.Map // goroutine中读写一致性， 原子性，用于全局变量
var key = "smap"

func main() {
	fmt.Println(smap.LoadOrStore(key, "value1"))
	fmt.Println(smap.Load(key))
	smap.Store(key, "value2")
	fmt.Println(smap.Load(key))
	smap.Delete(key)
	fmt.Println(smap.Load(key))
}
