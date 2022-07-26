package Algo

func strCmp(srcOri, dstOri map[byte]int) bool {
	src := make(map[byte]int, len(srcOri))
	dst := make(map[byte]int, len(dstOri))
	for k, v := range srcOri {
		src[k] = v
	}
	for k, v := range dstOri {
		dst[k] = v
	}
	for k, v := range src {
		if dst[k] == v {
			//delete(src, k)
			delete(dst, k)
		} else {
			return false
		}
	}
	if len(dst) != 0 {
		return false
	} else {
		return true
	}
}
func GroupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	strMaps := make([]map[byte]int, 0)
	for i := 0; i < len(strs); i++ {
		strMap := make(map[byte]int, 0)
		for _, c := range []byte(strs[i]) {
			strMap[c]++
		}
		strMaps = append(strMaps, strMap)
	}
	//fmt.Printf("%v\n", len(strMaps))
	for len(strs) != 0 {
		tmp := []string{strs[0]}
		strs = strs[1:]
		strMap := make(map[byte]int, len(strMaps[0]))
		for k, v := range strMaps[0] {
			strMap[k] = v
		}
		strMaps = strMaps[1:]
		//fmt.Printf("%v\n", strMap)
		for i := 0; i < len(strs); i++ {
			if strCmp(strMap, strMaps[i]) {
				tmp = append(tmp, strs[i])
				strs = append(strs[:i], strs[i+1:]...)
				strMaps = append(strMaps[:i], strMaps[i+1:]...)
				i--
			}
		}
		res = append(res, tmp)
	}
	return res
}
