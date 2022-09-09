package algo

func countRes(curMap map[byte][]int, c byte, preCnt [][2]int) (curCnt [][2]int) {
	if len(preCnt) == 0 {
		for k := range curMap[c] {
			curCnt = append(curCnt, [2]int{1, k})
		}
	} else {
		preRepeat, curRepeat := 0, 0
		iMap, iCnt := 0, 0
		for iMap < len(curMap[c]) {
			if preCnt[iCnt][1] < curMap[c][iMap] {
				if curRepeat > 0 {
					curCnt = append(curCnt, [2]int{preRepeat * curRepeat, curMap[c][iMap]})
					preRepeat, curRepeat = 0, 0
				}
				preRepeat += preCnt[iCnt][0]
				iCnt++
			} else {
				curRepeat++
				iMap++
			}
		}
	}
	return
}

func NumDistinct(s string, t string) (res int) {
	strMap := make(map[byte][]int)
	for k := range t {
		if _, ok := strMap[t[k]]; !ok {
			strMap[t[k]] = make([]int, 0)
		}
	}
	for k := range s {
		if _, ok := strMap[s[k]]; ok {
			strMap[s[k]] = append(strMap[s[k]], k)
		}
	}
	curCnt := make([][2]int, 0)
	for k := range t {
		curCnt = countRes(strMap, t[k], curCnt)
	}
	for k := range curCnt {
		res += curCnt[k][0]
	}
	return
}
