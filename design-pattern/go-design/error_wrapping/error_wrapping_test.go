package error_wrapping

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestAppError_Error(t *testing.T) {
	err := firstCall(10)
	fmt.Println(err)
	if err != nil {
		switch v := errors.Cause(err).(type) {
		case *AppError:
			fmt.Println("Custom App Error", v)
		default:
			fmt.Println(err)
		}
	}
}
