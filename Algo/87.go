package algo

func pickChar(s string, c byte) []int {
	res := make([]int, 0)
	for k := range s {
		if c == s[k] {
			res = append(res, k)
		}
	}
	return res
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) == 1 {
		return s1 == s2
	} else if len(s1) == 2 {
		return s1 == s2 || (s1[0] == s2[1] && s1[1] == s2[0])
	}
	/*from, to := 0, 0
	for k := range s1 {

	}*/
	return false
}
