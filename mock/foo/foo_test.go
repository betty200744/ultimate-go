package foo

import (
	"github.com/golang/mock/gomock"
	mock_foo "gobyexample/mock/foo/mock"
	"testing"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_foo.NewMockFoo(ctrl)
	m.
		EXPECT().
		Bar(1).
		AnyTimes()
}
