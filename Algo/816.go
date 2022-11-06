package algo

func groups(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	res := make([]string, 0)
	if s[0] != '0' || s[len(s)-1] != '0' {
		if s[0] == '0' {
			res = []string{s[:1] + "." + s[1:]}
		} else if s[len(s)-1] == '0' {
			res = []string{s}
		} else {
			res = append(res, s)
			for i := 1; i < len(s); i++ {
				res = append(res, s[:i]+"."+s[i:])
			}
		}
	}
	return res
}

func ambiguousCoordinates(s string) []string {
	res := make([]string, 0)
	for k := 2; k < len(s)-1; k++ {
		left := groups(s[1:k])
		if len(left) == 0 {
			continue
		}
		for _, l := range left {
			right := groups(s[k : len(s)-1])
			if len(right) == 0 {
				break
			}
			for _, r := range right {
				res = append(res, "("+l+", "+r+")")
			}
		}
	}
	return res
}
