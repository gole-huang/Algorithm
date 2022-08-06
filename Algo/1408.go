package Algo

import (
	"sort"
	"strings"
)

/*StringMatching
Given an array of string words. Return all strings in words which is substring of another word in any order.
String words[i] is substring of words[j],if can be obtained removing some characters to left and/or right side of words[j].
*/
type stringSort []string

func (s stringSort) Len() int {
	return len(s)
}
func (s stringSort) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
func (s stringSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func StringMatching(words []string) []string {
	sort.Sort(stringSort(words))
	res := make([]string, 0)
	for i := 0; i < len(words)-1; i++ {
		word := words[i]
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], word) {
				res = append(res, word)
				break
			}
		}
	}
	return res
}
