package waitgroup

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func getUrl(wg *sync.WaitGroup, url string, ch chan string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	b, _ := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	wg.Done()
	ch <- string(b)
}
func WaitGroup() {
	// WaitGroup: first: init wg
	// WaitGroup: then: before goroutine wg.add
	// WaitGroup: then: after goroutine wg.done
	// WaitGroup: last: main goroutine wg.wait
	resultCh := make(chan string)
	wg := &sync.WaitGroup{}
	urls := []string{"http://www.baidu.com", "https://gorm.io"}
	for _, url := range urls {
		wg.Add(1)
		go getUrl(wg, url, resultCh)
	}
	wg.Wait()
	for i := 0; i < len(urls); i++ {
		fmt.Println("result: ", <-resultCh)

	}
}
