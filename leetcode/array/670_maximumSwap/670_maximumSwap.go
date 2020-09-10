package _70_maximumSwap

import (
	"strconv"
	"strings"
)

func MaximumSwap(num int) int {
	arr := make([]string, 0)
	for _, i2 := range strconv.FormatInt(int64(num), 10) {
		arr = append(arr, string(i2))
	}

	max_index := -1
	max_value := ""
	l_index := -1
	r_index := -1
	// 右边最大的只能换左边比自己小的
	for i := len(arr) - 1; i >= 0; i-- {
		// 找当前最大的index
		if arr[i] > max_value {
			max_index = i
			max_value = arr[i]
			continue
		}
		// 找出最左边的比"当前最大的"小的， 注意没有比我小的就不会换
		if arr[i] < max_value {
			l_index = i
			r_index = max_index
		}
	}
	if l_index == -1 {
		return num
	}
	arr[l_index], arr[r_index] = arr[r_index], arr[l_index]
	r, _ := strconv.Atoi(strings.Join(arr, ""))
	return r
}
