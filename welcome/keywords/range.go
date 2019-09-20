package keywords

import "fmt"

func main() {
	// range, string , array, slice, map,
	r := "string"
	rr := []int{1, 2, 3}
	rr = append(rr, 1)
	rrr := make([]int, 0)
	rrr = append(rrr, 3)
	rrrr := map[string]string{"a": "a"}

	for key, value := range r {
		fmt.Println(key, value)
	}

	for _, value := range r {
		fmt.Println(value)
	}

	for key := range r {
		fmt.Println(key)
	}

	for key, value := range rr {
		fmt.Println(key, value)
	}

	for key, value := range rrr {
		fmt.Println(key, value)
	}

	for key, value := range rrrr {
		fmt.Println(key, value)
	}
}
