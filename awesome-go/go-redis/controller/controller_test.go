package controller

import (
	"context"
	"os"
	"testing"
	"time"
)

func init() {
	os.Setenv("ELASTIC_APM_SERVICE_NAME", "go-redis")
	os.Setenv("ELASTIC_APM_SERVER_URL", "http://127.0.0.1:8200")
}
func TestExampleClient(t *testing.T) {
	for i := 0; i < 5; i++ {
		ExampleClient(context.Background())
	}
	time.Sleep(time.Second * 2)
}
func TestExampleClient2(t *testing.T) {
	for i := 0; i < 5; i++ {
		ExampleClient(context.Background())
	}
	time.Sleep(time.Second * 2)
}
