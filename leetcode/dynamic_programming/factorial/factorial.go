package factorial

func Factorial(n int) int {
	if n > 0 {
		return Factorial(n-1) * n
	}
	return 1
}

func FactorialDP_Tabulation(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] * i
	}
	return dp[n]
}
func FactorialDP_Memoization(n int) int {
	dp := make([]int, n+1)
	if n == 0 {
		return 1
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = n * FactorialDP_Memoization(n-1)
	return dp[n]
}
