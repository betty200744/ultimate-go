package controller

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("ELASTIC_APM_SERVICE_NAME", "go-redis")
	os.Setenv("ELASTIC_APM_SERVER_URL", "http://127.0.0.1:8200")
}
func TestExampleClient(t *testing.T) {
	ExampleClient()
}
