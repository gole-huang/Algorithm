package Algo

import (
	"sort"
	"strings"
)

func resortStr(s string) string {
	result := make([]string, 0)
	diff, start := 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			if diff == 0 {
				start = i
			}
			diff++
		} else {
			diff--
			if diff == 0 {
				str := depStr(s[start : i+1])
				result = append(result, str)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(result)))
	return strings.Join(result, "")
}

func depStr(s string) string {
	if s == "10" {
		return s
	}
	//先找可分离的前缀，前缀为全1
	prefix := 1
	for prefix < (len(s)+1)/2 {
		if s[prefix] == '1' && s[len(s)-1-prefix] == '0' {
			prefix++
		} else {
			break
		}
	}
	if prefix == (len(s)+1)/2 {
		return s
	} else { //后缀有可分离的序列
		pad := 0
		for pad < len(s)-prefix {
			if s[prefix+pad] == '1' {
				pad++
			} else {
				break
			}
		}
		zero := 0
		for prefix+pad+zero < len(s) {
			zero++
			if s[prefix+pad+zero] == '1' {
				break
			}
		}
		fix := prefix + pad - zero
		if zero-1 < pad {
			//if s[len(s)-1-prefix] == '1'
			return strings.Repeat("1", prefix-1) + resortStr(s[prefix-1:len(s)-prefix+1]) + strings.Repeat("0", prefix-1)
		} else {
			return strings.Repeat("1", fix) + resortStr(s[fix:len(s)-fix]) + strings.Repeat("0", fix)
		}
	}
}

/*
MakeLargestSpecial
一个序列必定是a个1，b个0，c个1开始，d个0结尾。其中a>=b，a<=d。若a=b，则直接返回
如果b<c，则代表
*/
func MakeLargestSpecial(s string) string {
	return resortStr(s)
}
