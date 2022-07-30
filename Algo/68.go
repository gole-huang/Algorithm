package Algo

import "strings"

/*
FullJustify
给定一个单词数组words和一个长度maxWidth，重新排版单词，使其成为每行恰好有maxWidth个字符，且左右两端对齐的文本。
你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格' '填充，使得每行恰好有 maxWidth个字符。
要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
文本的最后一行应为左对齐，且单词之间不插入额外的空格。
*/
func FullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	curLen := 0
	var line string
	for i := 0; i < len(words); i++ {
		curLen += len(words[i])
		if curLen == maxWidth || curLen == maxWidth-1 {
			if curLen == maxWidth-1 {
				words[0] += " "
			}
			for j := 0; j < i; j++ {
				line += words[j] + " "
			}
			line += words[i]
			res = append(res, line)
			words = words[i+1:]
			line = ""
			curLen, i = 0, -1
			continue
		} else if curLen < maxWidth-1 {
			curLen++
			continue
		} else { //curLen > maxWidth
			curLen -= len(words[i])
			//i--
			spaces := maxWidth - curLen + i
			if i == 1 {
				line += words[i-1] + strings.Repeat(" ", spaces)
			} else {
				space := spaces / (i - 1)
				moreSpace := strings.Repeat(" ", space+1)
				lessSpace := strings.Repeat(" ", space)
				leftSpace := spaces - space*(i-1)
				for j := 0; j < leftSpace; j++ {
					line += words[j] + moreSpace
				}
				for j := leftSpace; j < i-1; j++ {
					line += words[j] + lessSpace
				}
				line += words[i-1]
			}
			res = append(res, line)
			words = words[i:]
			line = ""
			curLen, i = 0, -1
			continue
		}
	}
	if len(words) != 0 {
		for i := 0; i < len(words)-1; i++ {
			line += words[i] + " "
		}
		line += words[len(words)-1]
		line += strings.Repeat(" ", maxWidth-len(line))
		res = append(res, line)
	}
	return res
}
