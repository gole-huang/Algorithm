package algo

func UniqueLetterString(s string) (res int) {
	charSet := [26][3]int{{}}
	for k, v := range s {
		v -= 65
		charSet[v][2] = charSet[v][1]
		charSet[v][1] = charSet[v][0]
		charSet[v][0] = k
		switch {
		case charSet[v][1] == 0:
			if s[k] == s[0] {
				res += k
			}
		case charSet[v][2] == 0 && s[k] != s[0]:
			res += (charSet[v][0] - charSet[v][1]) * (charSet[v][1] - charSet[v][2] + 1)
		default:
			res += (charSet[v][0] - charSet[v][1]) * (charSet[v][1] - charSet[v][2])
		}
	}
	for k, v := range charSet {
		if v[0] != 0 {
			if v[1] == 0 && byte(k+65) != s[0] {
				res += (len(s) - v[0]) * (v[0] - v[1] + 1)
			} else {
				res += (len(s) - v[0]) * (v[0] - v[1])
			}
		} else if byte(k+65) == s[0] && v[0] == 0 {
			res += (len(s) - v[0]) * (v[0] - v[1] + 1)
		}
	}
	return
}
