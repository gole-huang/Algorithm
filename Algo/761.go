package Algo

/*
MakeLargestSpecial
一个序列必定是a个1，b个0，c个1开始，d个0结尾。其中a>=b，a<=d。若a=b，则直接返回
*/
func MakeLargestSpecial(s string) string {
	maxQueue, start := 0, 0
	for start < len(s) {
		if s[start] == '1' {
			start++
		} else {
			break
		}
	}
	head := start
	for start < len(s) {
		if s[start] == '0' {
			start++
		} else {
			break
		}
	}
	for i := start; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' {
				if maxQueue < i-start {
					maxQueue = i - start
				}
			}
		} else {
			if s[i-1] == '0' {
				start = i
			}
		}
	}
	var tail int
	diff := 0
	for i := start; i < len(s); i++ {
		if s[i] == '1' {
			diff++
		} else {
			diff--
			if diff == 0 {
				tail = i + 1
				break
			}
		}
	}

	for i := start - 1; i >= 0; i-- {
		if s[i] == '0' {
			diff++
		} else {
			diff--
			if diff == 0 {
				//交换数组
				tmpHead := s[:i]
				tmpStr := s[i:start]
				str := s[start:tail]
				tmpTail := s[tail:]
				s = tmpHead + str + tmpStr + tmpTail
				start = i
				tail -= start - i
			}
		}
	}
}
