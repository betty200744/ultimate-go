package myconfig

import (
	"encoding/json"
	"io/ioutil"
)

func Load(path string, config interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return Parse(buf, config)
}
func Parse(buf []byte, config interface{}) error {
	return json.Unmarshal(buf, config)
}
