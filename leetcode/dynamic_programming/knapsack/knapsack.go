package knapsack

import "fmt"

type Item struct {
	weight int
	value  int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// values []
// weights []
// W = 50
// solutions
func Knapsack(items []Item, W int) int {
	arr := make([][]int, len(items)+1)
	for i := range arr {
		arr[i] = make([]int, W+1)
	}
	for i := 0; i <= len(items); i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				arr[i][w] = 0
				continue
			}
			if items[i-1].weight > w {
				arr[i][w] = arr[i-1][w]
			} else {
				arr[i][w] = max(arr[i-1][w-items[i-1].weight]+items[i-1].value, arr[i-1][w])
			}
		}
	}
	fmt.Println(arr[len(items)])
	return arr[len(items)][W]
}
