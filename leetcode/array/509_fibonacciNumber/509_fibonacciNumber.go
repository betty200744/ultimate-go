package _09_fibonacciNumber

func FibonacciNumber(N int) int {
	if N <= 1 {
		return N
	}
	return FibonacciNumber(N-1) + FibonacciNumber(N-2)
}
func FibonacciNumberDP(N int) int {
	f := make([]int, N+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= N; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[N]
}
