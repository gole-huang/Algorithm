package Algo

/*
MakeLargestSpecial
首先寻找特殊序列中除了头部以外，最长的连续‘1’，并找到该连续‘1’开头的子序列，以s[start:tail]标记
从start-1往前追溯，每遇到完整一个子序列就交换位置，直到最靠前。
然后针对
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
	for start < len(s) {
		if s[start] == '0' {
			start++
		} else {
			break
		}
	}
	maxStart := 0
	for i := start; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' {
				if maxQueue < i-start {
					maxQueue = i - start
					maxStart = start
				}
			}
		} else {
			if s[i-1] == '0' {
				start = i
			}
		}
	}
	if maxQueue == 1 {
		return s
	}
	start = maxStart
	var tail int
	diff, i := 0, start
	for ; i < len(s); i++ {
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
	if i == len(s) {
		return s
	}
	i = start - 1
	for ; i >= 0; i-- {
		if s[i] == '0' {
			diff++
		} else {
			diff--
			if diff == 0 {
				tmpHead := s[:i]
				tmpStr := s[i:start]
				str := s[start:tail]
				tmpTail := s[tail:]
				if tmpStr < str {
					s = tmpHead + str + tmpStr + tmpTail
					tail -= start - i
					start = i
				} else if tail == len(s)-1 {
					return s
				} else {
					return s[:start] + MakeLargestSpecial(s[start:])
				}
			}
		}
	}
	return MakeLargestSpecial(s)
}
