package algo

func beautySum(s string) int {
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		dp[i] = dp[i-1]
		strCnt := [26]int{}
		strCnt[s[i]-97] = 1
		max, min := 1, 1
		for d := i - 1; d >= 0; d-- {
			if strCnt[s[d]-97] == max {
				max++
			}
			strCnt[s[d]-97]++
			if strCnt[s[d]-97] == 1 {
				min = 1
			} else if strCnt[s[d]-97] == min+1 {
				addMin := true
				for _, v := range strCnt {
					if v == min {
						addMin = false
						break
					}
				}
				if addMin {
					min++
				}
			}
			dp[i] += max - min
		}
	}
	return dp[len(dp)-1]
}
