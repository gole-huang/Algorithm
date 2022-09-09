package algo

func search(s, p string) int {
	for i := 0; i < len(s); i++ {
		fullMatch := true
		for j := 0; j < len(p); j++ {
			if i+j >= len(s) {
				return -1
			} else if p[j] == '?' || p[j] == s[i+j] {
				continue
			} else {
				fullMatch = false
				break
			}
		}
		if fullMatch {
			return i
		}
	}
	return -1
}

func IsMatch(s, p string) bool {
	if p == "" {
		return s == ""
	}
	match := make([]string, 0)
	start, end := 0, 0
	hasStar := false
	for i := 0; i < len(p); i++ {
		if p[i] != '*' {
			end++
		} else {
			hasStar = true
			if i < len(p)-1 && p[i+1] == '*' {
				continue
			} else {
				if start != end {
					match = append(match, p[start:end])
				}
				start = i + 1
				end = i + 1
			}
		}
	}
	if start != end {
		match = append(match, p[start:end])
	}
	if p[0] != '*' {
		if search(s, match[0]) != 0 {
			return false
		} else {
			s = s[len(match[0]):]
			match = match[1:]
		}
	}
	if len(match) > 0 {
		lastWordLength := len(match[len(match)-1]) // 最后一个单词长度
		if p[len(p)-1] != '*' {
			if len(s) < lastWordLength || search(s[len(s)-lastWordLength:], match[len(match)-1]) != 0 {
				return false
			} else {
				s = s[:len(s)-lastWordLength]
				match = match[:len(match)-1]
			}
		}
	}
	for i := 0; i < len(match); i++ {
		tmpRes := search(s, match[i])
		if tmpRes != -1 {
			s = s[tmpRes+len(match[i]):]
		} else {
			return false
		}
	}
	return s == "" || hasStar
}
