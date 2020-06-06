package pgtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Mapping struct {
	T map[string]interface{}
}

func (p Mapping) Value() (driver.Value, error) {
	j, _ := json.Marshal(p.T)
	return j, nil
}
func (p *Mapping) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i map[string]interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	p.T = i
	return nil
}
