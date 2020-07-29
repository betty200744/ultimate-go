package error_wrapping

import (
	"fmt"
	"github.com/pkg/errors"
)

// ---------------
// Wrapping Errors
// ---------------
type AppError struct {
	State int
}

func (a *AppError) Error() string {
	return fmt.Sprintf("App Error, State: %v", a.State)
}

func firstCall(i int) error {
	if err := secondCall(); err != nil {
		return errors.Wrapf(err, "%v, firstCall -> secondCall", i)
	}
	return nil
}
func secondCall() error {
	if err := thirdCall(); err != nil {
		return errors.Wrap(err, "secondCall -> thirdCall")
	}
	return nil
}
func thirdCall() error {
	return &AppError{State: 99}
}
