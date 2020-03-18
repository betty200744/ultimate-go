package main

import (
	"fmt"
	"sort"
)

type SkuItem struct {
	Id    int64  `json:"id"`
	Index int64  `json:"index"`
	Name  string `json:"name"`
}
type ById []SkuItem

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }

type ByIndex []SkuItem

func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIndex) Less(i, j int) bool { return a[i].Index < a[j].Index }

type ByOtherIds struct {
	Ids   []int64   `json:"ids"`
	Items []SkuItem `json:"items"`
}

func (a ByOtherIds) Len() int      { return len(a.Items) }
func (a ByOtherIds) Swap(i, j int) { a.Items[i], a.Items[j] = a.Items[j], a.Items[i] }
func (a ByOtherIds) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	skus := []SkuItem{
		{Id: 343, Index: 3, Name: "cc"},
		{Id: 341, Index: 1, Name: "aa"},
		{Id: 342, Index: 2, Name: "bb"},
	}
	fmt.Println("original: ", skus)
	sort.Sort(ById(skus))
	fmt.Println("sort by id: ", skus)
	sort.Sort(ByIndex(skus))
	fmt.Println("sort by index:", skus)
	otherIds := []int64{3, 1, 2}
	sortByOtherIds := ByOtherIds{
		Ids:   otherIds,
		Items: skus,
	}
	sort.Sort(sortByOtherIds)
}
