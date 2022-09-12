package algo

import "strconv"

func MaximumSwap(num int) int {
	str := []byte(strconv.Itoa(num))
	i, same := 0, 0
	for ; i < len(str)-1; i++ {
		if str[i] < str[i+1] {
			break
		} else if str[i] == str[i+1] {
			same++
		} else {
			same = 0
		}
	}
	if i == len(str)-1 {
		return num
	} else {
		max := i
		for j := i + 1; j < len(str); j++ {
			if str[max] <= str[j] {
				max = j
			}
		}
		if str[max] > str[0] {
			str[0], str[max] = str[max], str[0]
		} else {
			str[i-same], str[max] = str[max], str[i-same]
		}
		if newNum, err := strconv.Atoi(string(str)); err == nil {
			return newNum
		} else {
			return -1
		}
	}
}
