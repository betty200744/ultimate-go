package str

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
func Base64Decode(s string) string {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(b[:])
}
func IdsString2Slice(s string) []int64 {
	if s == "" {
		return []int64{}
	}
	arr := strings.Split(s, ",")
	n := len(arr)
	ids := make([]int64, 0, n)
	for _, s2 := range arr {
		i64, err := strconv.ParseInt(s2, 10, 64)
		if err == nil {
			ids = append(ids, i64)
		}
	}
	return ids
}
func Slice2IdsString(ids []int64) string {
	arr := make([]string, len(ids))
	for i, id := range ids {
		arr[i] = fmt.Sprintf("%d", id)
	}
	s := ""
	s = strings.Join(arr, ",")
	return s
}
