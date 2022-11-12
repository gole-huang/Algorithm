package algo

func customSortString(order string, s string) string {
	count := make([]int, len(order))
	str := make([]byte, len(s))
	ptr := 0
	for ks := range s {
		nomatch := true
		for ko := range order {
			if s[ks] == order[ko] {
				count[ko]++
				nomatch = false
				break
			}
		}
		if nomatch {
			str[ptr] = s[ks]
			ptr++
		}
	}
	for k := range order {
		for i := 0; i < count[k]; i++ {
			str[ptr] = order[k]
			ptr++
		}
	}
	return string(str)
}
