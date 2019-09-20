package main

import (
	"fmt"
	"github.com/grpc/grpc-go/status"
	"google.golang.org/grpc/codes"
)

func panic1() {
	panic("error in panic1")
}

func panic2() {
	panic("error in panic2")
}

func main() {

	defer func() {
		if errRecover := recover(); errRecover != nil {
			var err error
			isError, ok := errRecover.(error)
			if ok {
				err = status.Error(codes.Aborted, isError.Error())
			} else {
				err = status.Error(codes.Aborted, fmt.Sprintf("%v", errRecover))
			}
			fmt.Println(err)
		}
	}()
	panic("error in main")
	panic1()
	panic2()

}
