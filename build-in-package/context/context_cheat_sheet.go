package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// context.WithCancel, Manual cancel
	ctxWithCancel, cancelWithCancel := context.WithCancel(context.Background())
	cancelWithCancel()
	select {
	case a := <-time.After(time.Second):
		fmt.Println("case : time.After:", a)
	case d := <-ctxWithCancel.Done():
		fmt.Println("case: ctxWithCancel.Done:", ctxWithCancel.Err(), d) // context canceled
	}
	cancelWithCancel()
	fmt.Println("run cancel again!!! , but do nothing")

	// context.WithTimeout, Timeout Cancel
	ctxWithTimeout, cancelWithTimeout := context.WithTimeout(context.Background(), 50*time.Millisecond)
	select {
	case a := <-time.After(1 * time.Second):
		fmt.Println("case : time.After:", a)
	case d := <-ctxWithTimeout.Done():
		fmt.Println("case: ctxWithTimeout.Done:", ctxWithTimeout.Err(), d) //  context deadline exceeded
	}
	cancelWithTimeout()
	fmt.Println("run cancel!!!! , but do nothing")

	// context.WithDeadline, Deadline timeout
	t := time.Now().Add(time.Second * 2)
	ctxWithDeadline, cancelWithDeadline := context.WithDeadline(context.Background(), t)
	select {
	case a := <-time.After(10 * time.Second):
		fmt.Println("case : time.After:", a)
	case d := <-ctxWithDeadline.Done():
		fmt.Println("case: ctxWithDeadline.Done:", ctxWithDeadline.Err(), d) //  context deadline exceeded
	}
	cancelWithDeadline()
	fmt.Println("run cancel!!!! , but do nothing")

	// context.WithValue, put value to map[interface{}]interface{}
	ctxWithName := context.WithValue(context.Background(), "name", "betty")
	// context.Value, get value from ctx
	fmt.Println("get ctx name: ", ctxWithName.Value("name"))

}
