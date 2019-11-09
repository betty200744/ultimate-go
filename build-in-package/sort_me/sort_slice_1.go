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
	parents := make([]Parent, 0)
	ids := []string{"4", "1", "5", "3", "2"} //期望结果
	p1 := Parent{"3", "betty"}
	p2 := Parent{"1", "cc"}
	p3 := Parent{"2", "dd"}
	p4 := Parent{"4", "ee"}
	p5 := Parent{"5", "ff"}
	parents = append(parents, p1, p2, p3, p4, p5)
	fmt.Println("before: ", parents)

	// sort each Parent in the parents slice by Id
	for k := 0; k < len(ids); k++ {
		sort.Slice(parents, func(i, j int) bool {
			iIndex := SliceIndex(len(ids), func(k int) bool { return ids[k] == parents[i].id })
			jIndex := SliceIndex(len(ids), func(k int) bool { return ids[k] == parents[j].id })
			return iIndex < jIndex
		})
	}
	fmt.Println("after:", parents)
}
