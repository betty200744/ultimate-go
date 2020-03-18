package main

import (
	"context"
	"fmt"
	"time"
)

func Foo(ctx context.Context) {
	go Bar(ctx)
	time.Sleep(time.Second)
	fmt.Println("this is foo")

}
func Bar(ctx context.Context) {
	go Char(ctx)
	time.Sleep(time.Second)
	fmt.Println("this is bar")

}
func Char(ctx context.Context) {
	time.Sleep(time.Second)
	fmt.Println("this is char")
}
func main() {
	ctx := context.Background()
	go func() {
		Foo(ctx)
	}()
	fmt.Println("main")
	time.Sleep(time.Hour)
}
