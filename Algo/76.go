package algo

import (
	"strings"
)

func match(m map[byte][]int) bool {
	for _, v := range m {
		if v[0] > v[1] {
			return false
		}
	}
	return true
}

func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	} else if len(t) == 1 {
		if strings.Contains(s, t) {
			return t
		}
		return ""
	}
	// 初始化map
	strMap := map[byte][]int{}
	for k := range t {
		if _, ok := strMap[t[k]]; !ok {
			strMap[t[k]] = []int{1, 0}
		} else {
			strMap[t[k]][0]++
		}
	}
	// 双指针
	res := [2]int{}
	left, right := -1, 0
	for right < len(s) { // 右指针单独前移，取第一个完整子字符串
		if _, ok := strMap[s[right]]; ok {
			if left == -1 {
				left = right
			}
			strMap[s[right]][1]++
			if match(strMap) {
				if res[0] == res[1] || res[1]-res[0] > right-left {
					res = [2]int{left, right}
				}
				break
			}
		}
		right++
	}
	if left == -1 || right == len(s) {
		return ""
	}
	for left < right {
		if _, ok := strMap[s[left]]; ok {
			strMap[s[left]][1]--
		}
		left++
		if !match(strMap) {
			left--
			strMap[s[left]][1]++
			break
		}
	}
	res = [2]int{left, right}
	for {
		// 左右指针同时前移
		right++
		if right == len(s) {
			break
		}
		if _, ok := strMap[s[right]]; ok {
			strMap[s[right]][1]++
		}
		if _, ok := strMap[s[left]]; ok {
			strMap[s[left]][1]--
		}
		left++
		for match(strMap) {
			v, ok := strMap[s[left]]
			if !ok {
				left++
				continue
			} else if v[0] < v[1] {
				strMap[s[left]][1]--
				left++
				continue
			} else {
				break
			}
		}
		if match(strMap) && right-left < res[1]-res[0] {
			res = [2]int{left, right}
		}
	}
	return s[res[0] : res[1]+1]
}
