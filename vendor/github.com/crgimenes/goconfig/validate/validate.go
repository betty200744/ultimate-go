package validate

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/crgimenes/goconfig/structtag"
)

// Prefix is a string that would be placed at the beginning of the generated tags.
var Prefix string

// Usage is the function that is called when an error occurs.
var Usage func()

// Setup maps and variables
func Setup(tag string, tagDefault string) {
	structtag.Setup()
	structtag.Prefix = Prefix
	SetTag(tag)
	SetTagDefault(tagDefault)

	structtag.ParseMap[reflect.Int64] = reflectInt
	structtag.ParseMap[reflect.Int] = reflectInt
	structtag.ParseMap[reflect.Float64] = reflectFloat
	structtag.ParseMap[reflect.String] = reflectString
	structtag.ParseMap[reflect.Bool] = reflectBool
}

// SetTag set a new tag
func SetTag(tag string) {
	structtag.Tag = tag
}

// SetTagDefault set a new TagDefault to retorn default values
func SetTagDefault(tag string) {
	structtag.TagDefault = tag
}

// Parse configuration
func Parse(config interface{}) (err error) {
	err = structtag.Parse(config, "")
	return
}

// PrintDefaultsOutput changes the default output help string
var PrintDefaultsOutput string

func getValue(value *reflect.Value, datatype string) (ret string) {
	switch datatype {
	case "bool":
		if value.Bool() {
			ret = "true"
			return
		}
	case "string":
		ret = value.String()
		return
	case "int":
		ret = strconv.FormatInt(value.Int(), 10)
		return
	case "float64":
		f := value.Float()
		ret = strconv.FormatFloat(f, 'f', -1, 64)
		return
	}
	return
}

func reflectInt(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	req := field.Tag.Get("cfgRequired")
	valueStr := getValue(value, "int")
	if req == "true" && valueStr == "0" {
		err = fmt.Errorf("-%v is required", tag)
	}
	return
}

func reflectFloat(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	req := field.Tag.Get("cfgRequired")
	valueStr := getValue(value, "float64")
	if req == "true" && valueStr == "0" {
		err = fmt.Errorf("-%v is required", tag)
	}
	return
}

func reflectString(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	req := field.Tag.Get("cfgRequired")
	valueStr := getValue(value, "string")
	if req == "true" && valueStr == "" {
		err = fmt.Errorf("-%v is required", strings.ToLower(tag))
	}
	return
}

func reflectBool(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	return
}
