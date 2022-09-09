package algo

import (
	"strconv"
	"strings"
)

type logContent struct {
	id        int
	start     bool
	timeStamp int
}

func fmtLog(logs []string) []logContent {
	log := make([]logContent, len(logs))
	for l := range logs {
		content := strings.Split(logs[l], ":")
		log[l].id, _ = strconv.Atoi(content[0])
		log[l].start = content[1] == "start"
		log[l].timeStamp, _ = strconv.Atoi(content[2])
	}
	return log
}

func exclusiveTime(n int, logs []string) []int {
	times := make([]int, n)
	log := fmtLog(logs)
	stack := make([]int, 1)
	stack[0] = log[0].id
	curStamp := log[0].timeStamp
	for i := 1; i < len(log); i++ {
		if log[i].start {
			if len(stack) != 0 {
				times[stack[len(stack)-1]] += log[i].timeStamp - curStamp
			}
			stack = append(stack, log[i].id)
			curStamp = log[i].timeStamp
		} else {
			stack = stack[:len(stack)-1]
			log[i].timeStamp++
			times[log[i].id] += log[i].timeStamp - curStamp
			curStamp = log[i].timeStamp
		}
	}
	return times
}
