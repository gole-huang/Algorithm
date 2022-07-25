package Algo

func IsMatch(s string, p string) bool {
	if len(p) == 0 { // p匹配完成
		return len(s) == 0 //返回s是否匹配完成
	}
	i := 0
	for ; i < len(s); i++ {
		if len(p) <= i { // p完成匹配，s未完成
			return false
		} else if p[i] == '?' {
			continue
		} else if p[i] == '*' {
			otherStr := false
			for j := i; j <= len(s); j++ {
				k := 0
				for ; k < len(p)-i; k++ {
					if p[i+k] == '*' {
						continue
					} else {
						break
					}
				}
				if IsMatch(s[j:], p[i+k:]) {
					otherStr = true
					break
				}
			}
			return otherStr
		} else if s[i] != p[i] {
			return false
		}
	}
	for j := i; j < len(p); j++ { // s完成匹配，p把多余的*去掉
		if p[j] != '*' {
			return false
		}
	}
	return true
}