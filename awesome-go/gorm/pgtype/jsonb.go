package pgtype

import (
	"database/sql/driver"
	"errors"
)

type JSONB []byte

func (j JSONB) Value() (driver.Value, error) {
	return string(j), nil
}
func (j *JSONB) Scan(value interface{}) error {
	s, ok := value.([]byte)
	if !ok {
		errors.New("Scan source was not string")
	}
	*j = append((*j)[0:0], s...)

	return nil
}
