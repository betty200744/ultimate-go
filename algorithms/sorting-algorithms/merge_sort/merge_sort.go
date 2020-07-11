package main

import "fmt"

func merge(a []int, b []int) []int {
	r := make([]int, 0)
	i := 0
	j := 0
	// 注意此时的a, b都是已经排好序的
	// 每次都是a[i], b[j]比较，小的先放到r里面，此时"小的所在的数组"i加一位继续与b[j]对比
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r = append(r, a[i])
			i++
		} else {
			r = append(r, b[j])
			j++
		}
	}
	// 上面for循环结束后， 要么a还有数据， 要么b还有数据，即下面两个for循环只有一个会运行到
	// 因为剩下的数据是已经排好序的， 所有直接插入就可以
	for i < len(a) {
		r = append(r, a[i])
		i++
	}
	for j < len(b) {
		r = append(r, b[j])
		j++
	}
	return r
}
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var middle = len(arr) / 2
	var a = MergeSort(arr[:middle])
	var b = MergeSort(arr[middle:])
	return merge(a, b)
}

func main() {
	arr := []int{7, 2, 6, 9, 3, 5, 1, 8}
	after := MergeSort(arr)
	fmt.Println(after)
}
