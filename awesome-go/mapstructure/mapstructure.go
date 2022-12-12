package mapstructure

import (
	"github.com/mitchellh/mapstructure"
)

func MapStructureJsonTagDecode(input, output interface{}) {
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{Result: output, TagName: "json"})
	_ = decoder.Decode(input)
}
