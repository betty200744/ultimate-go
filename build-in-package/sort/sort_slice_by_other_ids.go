package main

import (
	"fmt"
	"sort"
)

type Parent struct {
	id   string
	name string
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func main() {
	//ids := []string{"4", "1", "5", "3", "2"} // 参考排序
	// 待排序列表
	parents := []Parent{
		{"3", "betty"},
		{"5", "cc"},
		{"2", "dd"},
		{"4", "ee"},
		{"1", "ff"},
	}

	// sort by id
	sort.Slice(parents, func(i, j int) bool {
		fmt.Println(i, j, parents)
		return parents[i].id < parents[j].id
	})
	fmt.Println("sort by id: ", parents)
	// sort each Parent in the parents slice by Id
	/*	for k := 0; k < len(ids); k++ {
		sort.Slice(parents, func(i, j int) bool {
			iIndex := SliceIndex(len(ids), func(k int) bool {
				//fmt.Println(k, ids[k], i, parents[i].id)
				return ids[k] == parents[i].id
			})
			jIndex := SliceIndex(len(ids), func(k int) bool { return ids[k] == parents[j].id })
			return iIndex < jIndex
		})
	}*/
}
