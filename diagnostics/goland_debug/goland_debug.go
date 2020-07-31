package goland_debug

import (
	"context"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

func FakeTraffic() {
	// Wait for the server to start
	time.Sleep(1 * time.Second)
	pages := []string{"/", "/login", "/logout", "/products", "/product/{productID}", "/basket", "/about"}
	activeConns := make(chan struct{}, 10)
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	i := int64(0)
	for {
		activeConns <- struct{}{}
		i++
		page := pages[rand.Intn(len(pages))]
		// We need to launch this using a closure function to
		// ensure that we capture the correct value for the
		// two parameters we need: page and i
		go func(p string, rid int64) {
			makeRequest(activeConns, c, p, rid)
		}(page, i)
	}
}
func FakeTraffic2() {
	// Wait for the server to start
	time.Sleep(1 * time.Second)
	pages := []string{"/", "/login", "/logout", "/products", "/product/{productID}", "/basket", "/about"}
	activeConns := make(chan struct{}, 10)
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	i := int64(0)
	for {
		activeConns <- struct{}{}
		i++
		page := pages[rand.Intn(len(pages))]
		// We need to launch this using a closure function to
		// ensure that we capture the correct value for the
		// two parameters we need: page and i
		go func(p string, rid int64) {
			labels := pprof.Labels("request", "automated", "page", p, "rid", strconv.Itoa(int(rid)))
			pprof.Do(context.Background(), labels, func(_ context.Context) {
				makeRequest(activeConns, c, p, rid)
			})
		}(page, i)
	}
}
func makeRequest(done chan struct{}, c *http.Client, page string, i int64) {
	defer func() {
		// Unblock the next request from the queue
		<-done
	}()
	page = strings.Replace(page, "{productID}", "abc-"+strconv.Itoa(int(i)), -1)
	r, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+page, nil)
	if err != nil {
		return
	}
	resp, err := c.Do(r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	time.Sleep(time.Duration(10+rand.Intn(40)) + time.Millisecond)
}
