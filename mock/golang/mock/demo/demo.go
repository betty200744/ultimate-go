package demo

import (
	"github.com/golang/mock/gomock"
	"os"
	"testing"
	mock_demo "ultimate-go/mock/golang/mock/mock"
)

//go:generate mockgen -source=demo.go -destination ../mock/demo_mock.go
type CloudConfigs interface {
	GetAll() (map[string]string, error)
	GetByKey(key string) string
}

func New(t *testing.T) CloudConfigs {
	if os.Getenv("ENVIRON") == "localhost" {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		return mock_demo.NewMockCloudConfigs(ctrl)
	} else {
		return new(CloudConfigsDemo)
	}
}

type CloudConfigsDemo struct {
}

func (cc *CloudConfigsDemo) GetAll() (map[string]string, error) {
	return map[string]string{"a": "a"}, nil
}
func (cc *CloudConfigsDemo) GetByKey(key string) string {
	mm := map[string]string{"a": "a", "b": "b"}
	return mm[key]
}
