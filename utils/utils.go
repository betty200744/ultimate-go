package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

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

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func Byte2String(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

func String2Byte(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}

func EscapeHttpHeader(input string) string {
	if len(input) == 0 {
		return input
	}
	withoutHttps := strings.TrimPrefix(strings.TrimSpace(input), "http://")
	return strings.TrimPrefix(withoutHttps, "https://")
}

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func Md5(input []byte) string {
	hash := md5.New()

	//Get the 16 bytes hash
	hash.Write(input)
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string

	return hex.EncodeToString(hashInBytes)
}

func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {

				return true
			}
		}
	}

	return false
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

func StructToMap(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("StructToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if name := fi.Name; name != "" {
			out[name] = v.Field(i).Interface()
		}
	}
	return out, nil
}
func ToMap(v interface{}) {

}
