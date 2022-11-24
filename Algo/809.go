package algo

type strNum struct {
	char  byte
	count int
}

func ExpressiveWords(s string, words []string) int {
	res := 0
	strNums := []strNum{{s[0], 1}}
	for i := 1; i < len(s); i++ {
		if strNums[len(strNums)-1].char == s[i] {
			strNums[len(strNums)-1].count++
		} else {
			strNums = append(strNums, strNum{s[i], 1})
		}
	}
	for _, w := range words {
		i, j := 0, 0
		for i < len(strNums) && j < len(w) {
			if strNums[i].char != w[j] {
				break
			} else {
				cnt := 1
				for j+cnt < len(w) && w[j+cnt] == w[j] {
					cnt++
				}
				if cnt == strNums[i].count || (cnt < strNums[i].count && strNums[i].count > 2) {
					i++
					j += cnt
				} else {
					break
				}
			}
		}
		if i == len(strNums) && j == len(w) {
			res++
		}
	}
	return res
}
