package knapsack

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
	N := len(items)
	arr := make([][]int, len(items)+1)
	for i := range arr {
		arr[i] = make([]int, W+1)
	}
	items = append([]Item{{weight: 0, value: 0}}, items...)
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if items[i].weight > w {
				arr[i][w] = arr[i-1][w]
			} else {
				arr[i][w] = max(arr[i-1][w-items[i].weight]+items[i].value, arr[i-1][w])
			}
		}
	}
	return arr[N][W]
}
