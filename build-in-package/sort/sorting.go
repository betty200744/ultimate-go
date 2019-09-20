package sort

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{1, 3, 4}
	strs := []string{"c", "a", "b"}
	sort.Ints(ints)
	sort.Strings(strs)
	fmt.Println(ints)
	fmt.Println(strs)
	fmt.Println("int is  sorted", sort.IntsAreSorted(ints))
	fmt.Println("string is  sorted", sort.StringsAreSorted(strs))
}
