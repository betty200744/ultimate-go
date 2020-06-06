package pgtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SendRuleStructSlice []*SendRuleStruct

func (p SendRuleStructSlice) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *SendRuleStructSlice) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	i := new([]*SendRuleStruct)
	if err := json.Unmarshal(source, i); err != nil {
		return err
	}

	*p = *i
	return nil
}
