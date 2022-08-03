package Algo

import (
	"fmt"
	"time"
)

func copyMap(srcMap map[string]int) map[string]int {
	dst := make(map[string]int)
	for k, v := range srcMap {
		dst[k] = v
	}
	return dst
}

func FindSubstring(s string, words []string) []int {
	res := make([]int, 0)
	lgh := len(words[0])
	fromFirst := (len(words) - 1) * lgh
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	for begin := 0; begin < len(words[0]); begin++ {
		str := s[begin:]
		needBack := 0
		wMap := copyMap(wordMap)
		t0 := time.Now()
		for p := 0; p < len(str); p += lgh {
			if p+lgh > len(str) {
				break
			}
			each := str[p : p+lgh]
			if _, ok := wMap[each]; !ok {
				wMap = copyMap(wordMap)
				if _, yes := wordMap[each]; yes {
					p -= lgh * needBack
					needBack = 0
				}
				continue
			}
			if wMap[each] == 1 {
				delete(wMap, each)
			} else {
				wMap[each]--
			}
			needBack++
			if len(wMap) == 0 {
				res = append(res, begin+p-fromFirst)
				p -= fromFirst
				needBack = 0
				wMap = copyMap(wordMap)
			}
		}
		escaped := time.Since(t0)
		fmt.Println("Used ", escaped)
	}
	return res
}
