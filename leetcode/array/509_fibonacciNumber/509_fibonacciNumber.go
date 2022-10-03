package _09_fibonacciNumber

func FibonacciNumber(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	if N > 1 {
		return FibonacciNumber(N-1) + FibonacciNumber(N-2)
	}
	return 0
}
