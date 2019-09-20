package keywords

import "fmt"

func finished() {
	fmt.Println("finished")
}

func main() {
	defer finished()
	fmt.Println("start")
	nums := []int{1, 3, 4, 5}
	for key, value := range nums {
		fmt.Println(key, value)
	}
}
