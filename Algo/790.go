package algo

func numTilings(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for k := 3; k <= n; k++ {
		for i := 0; i < k-2; i++ {
			dp[k] += dp[i] << 1
		}
		dp[k] = (dp[k] + dp[k-2] + dp[k-1]) % (1e9 + 7)
	}
	return dp[n]
}
