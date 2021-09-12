package user

import (
	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	mock_user "ultimate-go/mock/user/mock"
)

func TestIndex(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		// mockgen -source=user.go -destination=./mock/
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockIndex := mock_user.NewMockIndex(ctrl)
		mockIndex.EXPECT().Get("a").Return("haha")
		res := mockIndex.Get("a")
		assert.Equal(t, res, "haha")
	})
	t.Run("GetTwo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockIndex := mock_user.NewMockIndex(ctrl)
		mockIndex.EXPECT().GetTwo("", "").Return("res1", "res2")
		res1, res2 := mockIndex.GetTwo("", "")
		assert.Equal(t, res1, "res1")
		assert.Equal(t, res2, "res2")
	})
}
