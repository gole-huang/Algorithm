package algo

func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for km := range dp {
		dp[km] = make([]int, n)
		for kn := range dp[km] {
			if km == 0 || kn == 0 {
				dp[km][kn] = 1
			} else {
				dp[km][kn] = dp[km-1][kn] + dp[km][kn-1]
			}
		}
	}
	return dp[m-1][n-1]
}
