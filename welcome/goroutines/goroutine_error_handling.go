package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Res struct {
	Value int
	Err   error
}

func doExpensiveStuff(i int) (int, error) {
	var err error
	if rand.Intn(2) == 0 {
		err = fmt.Errorf("goroutine error returned")
		if err != nil {
			return 0, err
		}
	}
	return i, nil
}

func Do() {
	resCh := make(chan Res)
	for i := 1; i < 9; i++ {
		go func(i int, resCh chan Res) {
			result, err := doExpensiveStuff(i)
			resCh <- Res{result, err}
		}(i, resCh)
	}

	for i := 1; i < 9; i++ {
		r := <-resCh
		if r.Err != nil {
			log.Fatal(r.Err)
		} else {
			fmt.Println(r)
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	Do()
}
