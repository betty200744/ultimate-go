package main

import (
	"fmt"
	"sort"
)

type TwoSlices struct {
	main_slice  []int
	other_slice []int
}

type SortByOther TwoSlices

func (sbo SortByOther) Len() int {
	return len(sbo.main_slice)
}

func (sbo SortByOther) Swap(i, j int) {
	sbo.main_slice[i], sbo.main_slice[j] = sbo.main_slice[j], sbo.main_slice[i]
	sbo.other_slice[i], sbo.other_slice[j] = sbo.other_slice[j], sbo.other_slice[i]
}

func (sbo SortByOther) Less(i, j int) bool {
	return sbo.other_slice[i] < sbo.other_slice[j]
}

func main() {
	my_other_slice := []int{3, 5, 1, 2, 4}
	my_main_slice := []int{1, 2, 3, 4, 5} // sorted : {3,4,1,2,5}

	my_two_slices := TwoSlices{main_slice: my_main_slice, other_slice: my_other_slice}

	fmt.Println("Not sorted : ", my_two_slices.main_slice)

	sort.Sort(SortByOther(my_two_slices))
	fmt.Println("Sorted : ", my_two_slices.main_slice)

}
