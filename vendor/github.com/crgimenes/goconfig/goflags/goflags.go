package goflags

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"flag"

	"github.com/crgimenes/goconfig/structtag"
)

type parameterMeta struct {
	Kind  reflect.Kind
	Value interface{}
	Tag   string
}

var (
	parametersMetaMap map[*reflect.Value]parameterMeta
	visitedMap        map[string]*flag.Flag
	disableFags       bool

	// Preserve disable default values and get only visited parameters thus preserving the values passed in the structure, default false
	Preserve bool

	// Prefix is a string that would be placed at the beginning of the generated tags.
	Prefix string

	//Usage is a function to show the help, can be replaced by your own version.
	Usage func()
)

// Setup maps and variables
func Setup(tag, tagDefault, TagHelper string) {
	Usage = DefaultUsage
	parametersMetaMap = make(map[*reflect.Value]parameterMeta)
	visitedMap = make(map[string]*flag.Flag)

	structtag.Setup()
	structtag.Prefix = Prefix
	SetTag(tag)
	SetTagDefault(tagDefault)
	SetTagHelper(TagHelper)

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

// SetTagHelper set a new TagHelper
func SetTagHelper(tag string) {
	structtag.TagHelper = tag
}

// Parse configuration
func Parse(config interface{}) (err error) {
	if disableFags {
		return
	}

	flag.Usage = Usage
	err = structtag.Parse(config, "")
	if err != nil {
		return
	}

	flag.Parse()

	flag.Visit(loadVisit)

	for k, v := range parametersMetaMap {
		if _, ok := visitedMap[v.Tag]; !ok && Preserve {
			continue
		}

		switch v.Kind {
		case reflect.String:
			value := *v.Value.(*string)
			k.SetString(value)
		case reflect.Int:
			value := *v.Value.(*int)
			k.SetInt(int64(value))
		case reflect.Float64:
			value := *v.Value.(*float64)
			k.SetFloat(value)
		case reflect.Bool:
			value := *v.Value.(*bool)
			k.SetBool(value)
		}
	}

	disableFags = true
	return
}

// Reset maps caling setup function
func Reset() {
	disableFags = false
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Usage = nil

	structtag.Reset()
	Setup(structtag.Tag, structtag.TagDefault, structtag.TagHelper)
}

func loadVisit(f *flag.Flag) {
	visitedMap[f.Name] = f
}

func reflectInt(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	var aux int
	var defaltValue string
	var defaltValueInt int

	defaltValue = field.Tag.Get(structtag.TagDefault)
	usage := field.Tag.Get(structtag.TagHelper)

	if defaltValue != "" && defaltValue != "0" {
		defaltValueInt, err = strconv.Atoi(defaltValue)
		if err != nil {
			return
		}
	}

	meta := parameterMeta{}
	meta.Value = &aux
	meta.Tag = strings.ToLower(tag)
	meta.Kind = reflect.Int
	parametersMetaMap[value] = meta

	flag.IntVar(&aux, meta.Tag, defaltValueInt, usage)

	return
}

func reflectFloat(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	var aux float64
	var defaltValue string
	var defaltValueFloat float64

	defaltValue = field.Tag.Get(structtag.TagDefault)
	usage := field.Tag.Get(structtag.TagHelper)

	if defaltValue != "" && defaltValue != "0" {
		defaltValueFloat, err = strconv.ParseFloat(defaltValue, 64)
		if err != nil {
			return
		}
	}

	meta := parameterMeta{}
	meta.Value = &aux
	meta.Tag = strings.ToLower(tag)
	meta.Kind = reflect.Float64
	parametersMetaMap[value] = meta

	flag.Float64Var(&aux, meta.Tag, defaltValueFloat, usage)

	return
}

func reflectString(field *reflect.StructField, value *reflect.Value, tag string) (err error) {

	var aux, defaltValue string
	defaltValue = field.Tag.Get(structtag.TagDefault)
	usage := field.Tag.Get(structtag.TagHelper)

	meta := parameterMeta{}
	meta.Value = &aux
	meta.Tag = strings.ToLower(tag)
	meta.Kind = reflect.String
	parametersMetaMap[value] = meta

	flag.StringVar(&aux, meta.Tag, defaltValue, usage)

	return
}

func reflectBool(field *reflect.StructField, value *reflect.Value, tag string) (err error) {
	var aux bool
	defaltTag := field.Tag.Get(structtag.TagDefault)
	defaltValue := defaltTag == "true" || defaltTag == "t"
	usage := field.Tag.Get(structtag.TagHelper)

	meta := parameterMeta{}
	meta.Value = &aux
	meta.Tag = strings.ToLower(tag)
	meta.Kind = reflect.Bool
	parametersMetaMap[value] = meta

	flag.BoolVar(&aux, meta.Tag, defaltValue, usage)

	return
}

// PrintDefaults print the default help
func PrintDefaults() {
	flag.PrintDefaults()
}

// DefaultUsage is assigned for Usage function by default
func DefaultUsage() {
	fmt.Println("Usage")
	PrintDefaults()
}
