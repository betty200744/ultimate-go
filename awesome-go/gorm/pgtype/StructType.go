package pgtype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SendRuleStruct struct {
	SendLimit    int32 `json:"send_limit"`
	SendLimitNum int64 `json:"send_limit_num"`
}

func (p SendRuleStruct) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *SendRuleStruct) Scan(src interface{}) (err error) {
	source, ok := src.([]byte)
	if !ok {
		err = errors.New("Type assertion .([]byte) failed.")
		return
	}
	err = json.Unmarshal(source, p)
	return
}
