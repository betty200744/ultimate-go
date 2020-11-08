package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func RandomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(RandInt(65, 90)) != temp {
			temp = string(RandInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func Byte2String(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

func String2Byte(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}

func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}
	/*	switch d := v.(type) {
		case string:

		}*/
	return 0
}

// ToString converts a value to string.
func ToString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(int64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return fmt.Sprintf("%+v", value)
	}
}
func ToInt64(value interface{}) int64 {
	switch value := value.(type) {
	case string:
		n, _ := strconv.Atoi(value)
		return int64(n)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	default:
		return value.(int64)
	}
}
