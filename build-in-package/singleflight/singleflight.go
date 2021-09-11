package main

import (
	"net/http"
	"time"

	"github.com/smallnest/rpcx/log"
	"golang.org/x/sync/singleflight"
)

var requestGroup singleflight.Group

func HelloWorld(key string) (string, error) {
	log.Infof("call HelloWorld, key: %v", key)
	time.Sleep(time.Second)
	return key, nil
}
func fooHandle(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path
	v, err, shared := requestGroup.Do(key, func() (interface{}, error) {
		go func() {
			time.Sleep(500 * time.Millisecond)
			log.Infof("forget key %v", key)
			requestGroup.Forget(key)
		}()
		return HelloWorld("foo")
	})
	if err != nil {
		log.Errorf("foo handle err, path: %v,  shared:%v ", req.URL.Path, shared)
		return
	}
	log.Infof("foo handle ok, path: %v,  shared: %v ", req.URL.Path, shared)
	w.Write([]byte(v.(string)))
}

func barHandle(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Path
	ch := requestGroup.DoChan(key, func() (interface{}, error) {
		return HelloWorld(key)
	})

	var result singleflight.Result
	select {
	case result = <-ch:
	case <-time.After(500 * time.Millisecond):
		return
	}
	v, err, shared := result.Val, result.Err, result.Shared
	if err != nil {
		log.Errorf("bar handle err, path: %v,  shared:%v ", req.URL.Path, shared)
		return
	}
	log.Infof("bar handle ok, path: %v,  shared:%v ", req.URL.Path, shared)
	w.Write([]byte(v.(string)))
}

func main() {
	// benchmark test: echo "GET http://localhost:8080/foo" | vegeta attack -duration=1s -rate=10 | vegeta report
	http.HandleFunc("/foo", fooHandle)
	// benchmark test: echo "GET http://localhost:8080/bar" | vegeta attack -duration=1s -rate=10 | vegeta report
	http.HandleFunc("/bar", barHandle)
	http.ListenAndServe(":8080", nil)
}
