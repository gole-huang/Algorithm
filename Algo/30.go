package Algo

import "strings"

/*findStr
一、目标
给定s string，words []string，在s内找到所有特定子字符串s'，s'=words[i1]+words[i2]+...+words[in]。给出所有s'的位置
二、参数：
str string：目标字符串
words []string：目标词组，比原词组少了words[0]
cur int：标记原始s中words[0]的位置。
begin int：标记s中搜索words[0]的起始位置
三、变量
start int：标记连续词组出现的最左侧位置
checked []bool：标记已经检查过的词组
nocheck int:剩余未检查的数量
四、处理
从cur往左，每次推len(words[0])的字符，检查该子字符串是否在words内且未被检查过，是则标记已检查，nocheck--，否则标记当前位置为start
从cur+len(words[0])往右，每次推len(words[0])的字符，检查该子字符串是否在words内且未被检查过，是则标记已检查，nocheck--，否则返回失败。
若nocheck==0，则代表该子字符串满足要求，返回start
由主函数决定继续检查还是返回结果
*/
/*FindSubstring
目标：把words[0]从words中分离，每次从str找出word所在的位置cur，把str、words、cur传入函数。
若cur==-1，则代表不能找到该该词组，返回res；
若函数返回-1，则代表该位置的词组不能组成完整子串，应缩小为str[cur+1:]继续查找
若函数返回一个值，则代表该位置有一个最左侧开始的完整字串。把返回值加入res，缩小为str[start+1:]继续查找
*/

func findStr(str string, words []string, cur, begin int) (starts []int) {
	cur += begin
	start := cur
	nocheck := len(words)
	checked := make([]bool, nocheck)
	for ; nocheck != 0 && start >= len(words[0]); start -= len(words[0]) { // 左探
		word := str[start-len(words[0]) : start]
		w := 0
		for ; w < len(words); w++ {
			if !checked[w] && word == words[w] {
				nocheck--
				checked[w] = true
				break
			}
		}
		if w == len(words) { // 找不到可匹配的词，改右探
			break
		}
	}
	cur += len(words[0])
	for ; nocheck != 0 && cur <= len(str)-len(words[0]); cur += len(words[0]) { // 右探
		w := 0
		word := str[cur : cur+len(words[0])]
		for ; w < len(words); w++ {
			if !checked[w] && word == words[w] {
				nocheck--
				checked[w] = true
				break
			}
		}
		if w == len(words) { // 找不到可匹配的词，返回错误
			return []int{-1}
		}
	}
	if nocheck != 0 {
		return []int{-1}
	}
	starts = append(starts, start)
	wordsLen := (len(words) + 1) * len(words[0])
	for i := start; i <= len(str)-wordsLen-len(words[0]); i += len(words[0]) {
		if str[i:i+len(words[0])] != str[i+wordsLen:i+wordsLen+len(words[0])] {
			break
		} else {
			starts = append(starts, i+len(words[0]))
		}
	} // 开始寻找首尾重复的字符串
	return
}

func FindSubstring(s string, words []string) (res []int) {
	if len(words) == 1 {
		tmp := strings.Index(s, words[0])
		begin := 0
		for tmp != -1 {
			res = append(res, begin+tmp)
			begin += tmp + 1
			if begin >= len(s) {
				break
			}
			tmp = strings.Index(s[begin:], words[0])
		}
		return
	}
	word := words[0]
	words = words[1:]
	res = make([]int, 0)
	resMap := make(map[int]bool)
	begin := 0
	var cur int
	for {
		if begin >= len(s) {
			break
		}
		cur = strings.Index(s[begin:], word) // cur为压缩字符串的坐标
		if cur == -1 {
			break
		}
		starts := findStr(s, words, cur, begin) //	start为非压缩字符串的坐标
		if starts[0] == -1 {
			begin += cur + 1
		} else {
			for _, start := range starts {
				if !resMap[start] {
					res = append(res, start)
					resMap[start] = true
					begin = starts[0] + 1
				} else {
					begin = starts[len(starts)-1] + len(word) + 1
				}
			}
		}
	}
	return
}
