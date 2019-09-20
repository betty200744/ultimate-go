package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	urls := []string{"http://www.sina.com", "http://www.baidu.com", "http://www.v2ex.com", "http://www.v2ex.com"}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
			fmt.Println("this is url: ", url)
		}(url)
	}
	fmt.Println("before wait")
	wg.Wait()
	fmt.Println("after wait")
}
