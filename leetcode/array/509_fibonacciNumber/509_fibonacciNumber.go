package _09_fibonacciNumber

func FibonacciNumber(N int) int {
	if N <= 1 {
		return N
	}
	return FibonacciNumber(N-1) + FibonacciNumber(N-2)
}

// the best, Bottom-up DP Algorithm
func Fib_DP_Tabulation(N int) int {
	arr := make([]int, N+1)
	if N < 1 {
		arr[0] = 0
	} else {
		arr[0] = 0
		arr[1] = 1
	}
	for i := 2; i <= N; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[N]
}

// bad
func Fib_DP_memoized(N int) int {
	arr := make([]int, N+1)
	if arr[N] != 0 {
		return arr[N]
	}
	if N <= 1 {
		return N
	} else {
		arr[N] = Fib_DP_memoized(N-1) + Fib_DP_memoized(N-2)
	}
	return arr[N]
}
