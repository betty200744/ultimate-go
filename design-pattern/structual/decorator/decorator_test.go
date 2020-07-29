package decorator

import (
	"fmt"
	"testing"
)

func TestRepository_Fetch(t *testing.T) {
	repository := &Repository{}
	// Decorator func for retry n times when get error
	_ = RetryWithParams(func() error {
		data, errFetch := repository.Fetch()
		fmt.Println(data)
		return errFetch
	}, 2)
}
