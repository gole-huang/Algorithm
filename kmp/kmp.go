package kmp

func compStr(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func buildNext(p string) []int {
	next := make([]int, len(p))
	next[0] = -1
	for i := 1; i < len(next); i++ {
		for l := i / 2; l > 0; l-- {
			if compStr(p[:l], p[i-l:i]) {
				next[i] = l
				break
			}
		}
	}
	return next
}

func SubString(s, p string) int {
	i, j := 0, 0
	next := buildNext(p)
	for i < len(s) && j < len(p) {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(p) {
		return i - j
	} else {
		return -1
	}
}
