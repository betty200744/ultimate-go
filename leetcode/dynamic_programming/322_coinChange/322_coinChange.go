package _22_coinChange

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CoinChange(coins []int, amount int) int {
	arr := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		arr[i] = amount + 1
	}
	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				arr[i] = min(arr[i], arr[i-coins[j]]+1)
			}
		}
	}
	if arr[amount] <= amount {
		return arr[amount]
	} else {
		return -1
	}
}
