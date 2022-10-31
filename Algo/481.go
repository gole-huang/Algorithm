package algo

func MagicalString(n int) int {
	s := make([]int, n+1)
	s[0], s[1], s[2] = 1, 2, 2
	res, cnt, i, val := 0, 0, 0, 1
	for ; i < n; cnt++ {
		for j := 0; j < s[cnt]; j++ {
			s[i] = val
			i++
		}
		if cnt%2 == 0 {
			val = 2
			res += s[cnt]
		} else {
			val = 1
		}
	}
	if i > n {
		for d := 1; d <= i-n; d++ {
			if s[i-d] == 1 {
				res--
			}
		}
	}
	return res
}
