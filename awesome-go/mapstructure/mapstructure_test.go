package mapstructure

import (
	"testing"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/assert.v1"
)

type Person struct {
	Id       int               `json:"id"`
	Name     string            `json:"name"`
	Channels []string          `json:"channels"`
	UserCode string            `json:"user_code"`
	Extra    map[string]string `json:"extra"`
}

func Test_Decode(t *testing.T) {
	input1 := map[string]interface{}{"id": 1, "name": "betty", "channels": []string{"c1", "c2"}, "extra": map[string]string{"a": "a"}}
	output := &Person{}
	mapstructure.Decode(input1, output)
	assert.Equal(t, input1["id"].(int), output.Id)
	assert.Equal(t, input1["extra"].(map[string]string)["a"], output.Extra["a"])
}
func Test_DecodeArray(t *testing.T) {
	input1 := map[string]interface{}{"id": 1, "name": "betty", "channels": []string{"c1", "c2"}}
	input2 := map[string]interface{}{"id": 1, "name": "betty", "channels": []string{"c1", "c2"}}
	inputs := []map[string]interface{}{input1, input2}
	output := make([]*Person, 0)
	mapstructure.Decode(inputs, &output)
	assert.Equal(t, input1["id"].(int), output[0].Id)
}
func TestMapStructureJsonTagDecode(t *testing.T) {
	input1 := map[string]interface{}{"id": 1, "name": "betty", "channels": []string{"c1", "c2"}, "user_code": "code", "extra": map[string]string{"a": "a"}}
	output := &Person{}
	MapStructureJsonTagDecode(input1, output)
	assert.Equal(t, input1["user_code"].(string), output.UserCode)
	assert.Equal(t, input1["extra"].(map[string]string)["a"], output.Extra["a"])
}
