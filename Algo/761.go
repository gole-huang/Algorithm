package Algo

func resortStr(ch []byte, start int) int {
	if len(ch) < 6 { // 不用排，没法动
		return start
	}
	var tail int
	diff, i := 0, start
	for ; i < len(ch); i++ {
		if ch[i] == '1' {
			diff++
		} else {
			diff--
			if diff < 0 {
				return start
			} else if diff == 0 {
				tail = i + 1
				break
			}
		}
	}
	if i == len(ch) { // 往后非子字符串，没法动
		return start
	}
	i = start - 1
	for ; i >= 0; i-- {
		if ch[i] == '0' {
			diff++
		} else {
			diff--
			if diff < 0 {
				break
			} else if diff == 0 {
				// tmpHead, tmpStr, str, tmpTail := ch[:i], ch[i:start], ch[start:tail], ch[tail:]
				tmpStr, str := ch[i:start], ch[start:tail]
				if string(tmpStr) < string(str) {
					// ch = []byte(string(tmpHead) + string(str) + string(tmpStr) + string(tmpTail)) //无法使形参修改反馈到实参
					xchg := make([]byte, len(str))
					copy(xchg, str)
					xchg = append(xchg, tmpStr...)
					copy(ch[i:tail], xchg)
					tail -= start - i
					start = i
					for start >= 0 && ch[start] == '1' {
						start--
					}
				} else {
					break
				}
			}
		}
	}
	return start
}

func MakeLargestSpecial(s string) string {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		if str[i] == '1' {
			i = resortStr(str, i)
		}
	}
	return string(str)
}
