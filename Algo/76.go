package Algo

func checkMap(m map[byte][2]int) bool {
	for _, v := range m {
		if v[0] > v[1] {
			return false
		}
	}
	return true
}

/*
	func replaceChar(pre, cur byte, m map[byte][2]int) bool {
		m[pre] = [2]int{m[pre][0], m[pre][1] - 1}
		m[cur] = [2]int{m[cur][0], m[cur][1] + 1}
		return checkMap(m)
	}
*/
func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	} else if len(t) == 1 {
		for k := range s {
			if t[0] == s[k] {
				return t
			}
		}
		return ""
	}
	tMap := map[byte][2]int{} //初始化map
	for k := range t {
		if _, ok := tMap[t[k]]; !ok {
			tMap[t[k]] = [2]int{1, 0}
		} else {
			tMap[t[k]] = [2]int{tMap[t[k]][0] + 1, tMap[t[k]][1]}
		}
	}
	res := [2]int{}
	left, right := -1, 0 //双指针
	loop := false
	for right < len(s) {
		if loop { //左指针前移，遇到重复字符，一直前移到map[0]
			for left < right {
				if _, ok := tMap[s[left]]; ok {
					tMap[s[left]] = [2]int{tMap[s[left]][0], tMap[s[left]][1] - 1}
					if !checkMap(tMap) {
						break
					} else {
						if res[1]-res[0] > right-left {
							res = [2]int{left, right}
						}
					}
				}
				left++
			}
		}
		if _, ok := tMap[s[right]]; ok {
			if left == -1 {
				left = right
			}
			tMap[s[right]] = [2]int{tMap[s[right]][0], tMap[s[right]][1] + 1}
			if checkMap(tMap) {
				loop = true
				if res[0] == res[1] || res[1]-res[0] > right-left {
					res = [2]int{left, right + 1}
				}
			}
		}
	}
	if res[0] == res[1] {
		return ""
	} else {
		return s[res[0] : res[1]+1]
	}
}
