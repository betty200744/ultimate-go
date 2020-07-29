package decorator

import (
	"fmt"
	"github.com/avast/retry-go"
)

// Data returned by fetch
type Data map[string]string

// Repository of data
type Repository struct{}

// Fetch fetches data
func (r *Repository) Fetch() (Data, error) {
	data := Data{
		"user":     "root",
		"password": "swordfish",
	}
	fmt.Printf("Repository fetched data successfully: %v\n", data)
	return data, nil
}

func RetryWithParams(fn func() error, attempts uint) error {
	err := retry.Do(fn,
		retry.Attempts(attempts),
		retry.LastErrorOnly(true),
		retry.OnRetry(func(n uint, err error) {
		}),
	)
	return err
}
