package sync

import (
	"fmt"
	"sync"
	"testing"
)

var (
	clientConnMap sync.Map
)

// curd, Store, Load, LoadOrStore,  Delete, LoadAndDelete, Range
func Test_Map(t *testing.T) {
	clientConnMap.Store("client1", "1.1.1.1")
	v, ok := clientConnMap.Load("client1")
	fmt.Println("Load: ", v, ok)
	v, loaded := clientConnMap.LoadOrStore("client2", "2.2.2.2")
	fmt.Println("loadOrStore: ", v, loaded)
	clientConnMap.Range(func(key, value interface{}) bool {
		fmt.Println(fmt.Sprintf("Range key: %v, value: %v", key, value))
		return true
	})
	clientConnMap.Delete("client1")
}
