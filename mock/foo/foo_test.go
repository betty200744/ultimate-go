package foo

import (
	"github.com/golang/mock/gomock"
	"testing"
	mock_foo "ultimate-go/mock/foo/mock"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_foo.NewMockFoo(ctrl)
	m.EXPECT().Bar(1).AnyTimes()
}
