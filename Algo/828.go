package Algo

func UniqueLetterString(s string) int {
	strMap := make([]map[byte]int, 0)
	arr, cur := 0, 0
	for i := range s {
		strMap = append(strMap, map[byte]int{s[i]: 1})
		cur++
		useless := 0
		for d := len(strMap) - 2; d >= 0; d-- {
			if len(strMap[d]) < 26 {
				if _, ok := strMap[d][s[i]]; !ok {
					cur++
					strMap[d][s[i]] = 1
				} else if strMap[d][s[i]] == 2 {
					break
				} else if strMap[d][s[i]] == 1 {
					cur--
					strMap[d][s[i]]++
				}
			} else {
				if strMap[d][s[i]] == 2 {
					break
				} else if strMap[d][s[i]] == 1 {
					cur--
					strMap[d][s[i]]++
				}
				if d <= len(strMap)-52 {
					allMatchTwice := true
					for _, v := range strMap[d] {
						if v == 1 {
							allMatchTwice = false
							break
						}
					}
					if allMatchTwice {
						useless = d
					}
				}
			}
		}
		arr += cur
		strMap = strMap[useless:]
	}
	return arr
}
